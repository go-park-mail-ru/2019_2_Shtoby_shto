package main

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/user"
	handler "2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/security"
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	deployEnvVar = "DEPLOYAPI"
)

var (
	initFlag = flag.Bool("initial start", false, "Check your service")
	httpAddr = flag.String("address", ":8080", "HTTP listen address")
)

var (
	handlerService  handler.Handler
	securityService security.Security
	userService     user.HandlerUserService
	photoService    photo.HandlerPhotoService
	dbService       database.InitDBManager
)

func main() {
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	config.InitConfig(logger)

	conf := config.GetInstance()

	// Нужно вообще убрать эту тему с флагами
	*httpAddr = ":" + strconv.Itoa(conf.Port)
	logger.Println("API Url:", *httpAddr)

	dbService = database.Init()
	db, err := dbService.DbConnect("postgres", conf.DbConfig)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
	dm := database.NewDataManager(db)
	defer dm.CloseConnection()

	e := echo.New()
	initService(e, dm, conf)
	newServer(e, logger)
	if *initFlag {
		return
	}

	//TODO::great shutdown
	switch conf.Port {
	case 443:
		if err := e.StartTLS(*httpAddr, "keys/server.crt", "keys/server.key"); err != http.ErrServerClosed {
			logger.Fatalf("HTTPS server ListenAndServe: %v", err)
		}
	default:
		if err := e.Start(*httpAddr); err != http.ErrServerClosed {
			logger.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}
}

func newServer(e *echo.Echo, logger *log.Logger) {
	//router := route.NewRouterService(securityService)
	logger.Println("serving on", *httpAddr)

	apiURL := config.GetInstance().FrontendURL
	e.Use(middleware.Logger(), middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{apiURL},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//e.POST("/login", s.LoginEcho)
	//e.GET("/user", s.CheckSessionEcho(s.ImageSecurityEcho))

	e.Server = &http.Server{
		Addr: *httpAddr,
		//Handler:        router,
		ErrorLog:       logger,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func initService(e *echo.Echo, db database.IDataManager, conf *config.Config) {
	sessionService := security.NewSessionManager(conf.RedisConfig, conf.RedisPass, conf.RedisDbNumber)
	userService = user.CreateInstance(db)
	photoService = photo.CreateInstance(db)
	//transportService = handler.CreateInstance(sessionService)
	securityService = security.CreateInstance(sessionService, userService, photoService)
	user.NewUserHandler(e, userService)
	photo.NewPhotoHandler(e, photoService, userService)

}
