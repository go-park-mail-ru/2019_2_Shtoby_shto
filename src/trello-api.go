package main

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/board"
	"2019_2_Shtoby_shto/src/dicts/boardUsers"
	"2019_2_Shtoby_shto/src/dicts/card"
	"2019_2_Shtoby_shto/src/dicts/cardGroup"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/task"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/initDB"
	"2019_2_Shtoby_shto/src/security"
	"context"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	// TODO::"github.com/microcosm-cc/bluemonday"
	//"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	initFlag = flag.Bool("initial start", false, "Check your service")
)

var (
	securityService   security.HandlerSecurity
	userService       user.HandlerUserService
	photoService      photo.HandlerPhotoService
	boardService      board.HandlerBoardService
	boardUsersService boardUsers.HandlerBoardUsersService
	cardUsersService  сardUsers.HandlerCardUsersService
	cardService       card.HandlerCardService
	cardGroupService  cardGroup.HandlerCardGroupService
	taskService       task.HandlerTaskService
	dbService         initDB.InitDBManager
)

func main() {
	flag.Parse()
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := config.InitConfig(); err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}

	conf := config.GetInstance()

	httpAddr := ":" + strconv.Itoa(conf.Port)
	e.Logger.Info("API Url:", httpAddr)

	dbService = initDB.Init()
	db, err := dbService.DbConnect("postgres", conf.DbConfig)
	if err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}
	dm := database.NewDataManager(db)
	defer dm.CloseConnection()

	e.Logger.SetLevel(echoLog.INFO)
	initService(e, dm, conf)
	newServer(e, httpAddr)
	if *initFlag {
		return
	}

	// great shutdown
	go func() {
		switch conf.Port {
		case 443:
			if err := e.StartTLS(httpAddr, "keys/server.crt", "keys/server.key"); err != http.ErrServerClosed {
				e.Logger.Fatalf("HTTPS server ListenAndServe: %v", err)
			}
		default:
			if err := e.Start(httpAddr); err != http.ErrServerClosed {
				e.Logger.Fatalf("HTTP server ListenAndServe: %v", err)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func newServer(e *echo.Echo, httpAddr string) {
	e.Logger.Info("serving on", httpAddr)

	apiURL := config.GetInstance().FrontendURL
	// securityService.CheckSession,
	e.Use(middleware.Logger(), middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{apiURL},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}), securityService.CheckSession)

	e.Server = &http.Server{
		Addr:           httpAddr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func initService(e *echo.Echo, db database.IDataManager, conf *config.Config) {
	sessionService := security.NewSessionManager(conf.RedisConfig, conf.RedisPass, conf.RedisDbNumber)
	userService = user.CreateInstance(db)
	photoService = photo.CreateInstance(db)
	boardService = board.CreateInstance(db)
	boardUsersService = boardUsers.CreateInstance(db)
	cardUsersService = сardUsers.CreateInstance(db)
	cardService = card.CreateInstance(db)
	cardGroupService = cardGroup.CreateInstance(db)
	taskService = task.CreateInstance(db)
	securityService = security.CreateInstance(sessionService)
	user.NewUserHandler(e, userService, boardUsersService, cardUsersService, securityService)
	photo.NewPhotoHandler(e, photoService, userService, securityService)
	board.NewBoardHandler(e, userService, boardService, boardUsersService, cardService, securityService)
	card.NewCardHandler(e, userService, cardService, cardUsersService, taskService, securityService)
	cardGroup.NewCardGroupHandler(e, cardGroupService, securityService)
	task.NewTaskHandler(e, userService, taskService, securityService)
}
