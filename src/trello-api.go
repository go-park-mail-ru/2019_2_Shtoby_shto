package main

import (
	"2019_2_Shtoby_shto/file_service/file"
	"2019_2_Shtoby_shto/session_service/session"
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/board"
	"2019_2_Shtoby_shto/src/dicts/boardUsers"
	"2019_2_Shtoby_shto/src/dicts/card"
	"2019_2_Shtoby_shto/src/dicts/cardGroup"
	"2019_2_Shtoby_shto/src/dicts/cardTags"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/checkList"
	"2019_2_Shtoby_shto/src/dicts/comment"
	"2019_2_Shtoby_shto/src/dicts/hub"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/tag"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/initDB"
	"2019_2_Shtoby_shto/src/metric"
	"2019_2_Shtoby_shto/src/security"
	"context"
	"errors"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	// TODO::"github.com/microcosm-cc/bluemonday"
)

var (
	securityService   security.HandlerSecurity
	sessionService    security.SessionHandler
	userService       user.HandlerUserService
	photoService      photo.HandlerPhotoService
	boardService      board.HandlerBoardService
	boardUsersService boardUsers.HandlerBoardUsersService
	cardUsersService  сardUsers.HandlerCardUsersService
	cardService       card.HandlerCardService
	cardGroupService  cardGroup.HandlerCardGroupService
	commentService    comment.HandlerCommentService
	checkListService  checkList.HandlerCheckListService
	tagService        tag.HandlerTagService
	cardTagsService   cardTags.HandlerCardTagsService
)

func main() {
	flag.Parse()
	e := echo.New()

	e.GET("/swagger/*", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusPermanentRedirect, "https://app.swaggerhub.com/apis/aleksandrkhoroshenin/trello-api/4.0")
	})

	e.POST("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Register prometheus metrics
	metric.RegisterAccessHitsMetric("api_service")

	if err := config.InitConfig(); err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}

	conf := config.GetInstance()

	sessService, err := ConnectGRPC(conf.SecurityURL, "security_service")
	if err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}
	securityClient := session.NewSecurityClient(sessService)

	fileService, err := ConnectGRPC(conf.FileLoaderURL, "file_service")
	if err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}
	fileLoaderClient := file.NewIFileLoaderManagerClient(fileService)

	httpAddr := ":" + strconv.Itoa(conf.Port)
	e.Logger.Info("API Url:", httpAddr)

	dbService := initDB.Init()
	db, err := dbService.DbConnect("postgres", conf.DbConfig)
	if err != nil {
		e.Logger.Error(err)
		os.Exit(3)
	}
	dm := database.NewDataManager(db)
	defer dm.CloseConnection()

	e.Logger.SetLevel(echoLog.DEBUG)
	InitServices(e, dm, conf, securityClient, fileLoaderClient)
	newServer(e, httpAddr)

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
	e.Use(
		middleware.Logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{apiURL},
			AllowCredentials: true,
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
			ExposeHeaders:    []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
		}),
		securityService.CheckSession,
		AccessHitsMiddleware,
		checkCSRF)

	e.Server = &http.Server{
		Addr:           httpAddr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func checkCSRF(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		csrfRequest := ctx.Request().Header.Get(echo.HeaderXCSRFToken)
		csrfCurrent := ctx.Get("csrf_token")
		if ctx.Get("not_security") == "done" {
			return h(ctx)
		}
		if csrfRequest != csrfCurrent {
			return errors.New("csrf error")
		}
		return h(ctx)
	}
}

func InitServices(e *echo.Echo, db database.IDataManager, conf *config.Config, sessService session.SecurityClient, fileService file.IFileLoaderManagerClient) {
	sessionService = security.NewSessionManager(&sessService)
	userService = user.CreateInstance(db)
	photoService = photo.CreateInstance(db, conf, fileService)
	boardService = board.CreateInstance(db)
	boardUsersService = boardUsers.CreateInstance(db)
	cardUsersService = сardUsers.CreateInstance(db)
	cardService = card.CreateInstance(db, fileService)
	cardGroupService = cardGroup.CreateInstance(db)
	commentService = comment.CreateInstance(db)
	checkListService = checkList.CreateInstance(db)
	tagService = tag.CreateInstance(db)
	cardTagsService = cardTags.CreateInstance(db)
	securityService = security.CreateInstance(sessionService)
	user.NewUserHandler(e, userService, boardUsersService, cardUsersService, securityService)
	photo.NewPhotoHandler(e, photoService, userService, securityService)
	board.NewBoardHandler(e, userService, boardService, boardUsersService, cardService, cardUsersService, cardGroupService, tagService, cardTagsService, commentService, checkListService, securityService)
	card.NewCardHandler(e, userService, cardService, cardUsersService, tagService, cardTagsService, commentService, checkListService, securityService)
	cardGroup.NewCardGroupHandler(e, cardGroupService, securityService)
	comment.NewCommentHandler(e, userService, commentService, securityService)
	checkList.NewCheckListHandler(e, userService, checkListService, securityService)
	tag.NewTagHandler(e, userService, tagService, cardTagsService, securityService)
	// register ws channel
	h := hub.NewHub(cardUsersService)
	go h.Run()
	hub.NewWsHandler(e, h)
}

func ConnectGRPC(addr string, name string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	if err != nil {
		log.Error("Can't connect to security service:", err)
		return nil, err
	}
	return conn, nil
}

func AccessHitsMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		start := time.Now()
		log.Info(start)

		// Write hits metric
		if metric.AccessHits != nil {
			metric.AccessHits.With(prometheus.Labels{
				"path":        ctx.Request().URL.Path,
				"method":      ctx.Request().Method,
				"status_code": strconv.Itoa(ctx.Response().Status),
			}).Inc()
		}

		log.Info("Finish with status code ",
			" status_code: ", ctx.Response().Status,
			" work_time:", time.Since(start))
		return h(ctx)
	}
}
