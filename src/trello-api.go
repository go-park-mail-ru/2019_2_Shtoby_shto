package main

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/board"
	"2019_2_Shtoby_shto/src/dicts/boardUsers"
	"2019_2_Shtoby_shto/src/dicts/card"
	"2019_2_Shtoby_shto/src/dicts/cardGroup"
	"2019_2_Shtoby_shto/src/dicts/cardTags"
	сardUsers "2019_2_Shtoby_shto/src/dicts/cardUsers"
	"2019_2_Shtoby_shto/src/dicts/comment"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/tag"
	"2019_2_Shtoby_shto/src/dicts/user"
	"2019_2_Shtoby_shto/src/fileLoader"
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
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	// TODO::"github.com/microcosm-cc/bluemonday"
	//"github.com/prometheus/client_golang/prometheus"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	flag.Parse()
	e := echo.New()

	e.GET("/swagger/*", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusPermanentRedirect, "https://app.swaggerhub.com/apis/aleksandrkhoroshenin/trello-api/4.0")
	})

	e.POST("/api/v1/query", echo.WrapHandler(promhttp.Handler()))

	// Register prometheus metrics
	metric.RegisterAccessHitsMetric("api_service")

	if err := config.InitConfig(); err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}

	conf := config.GetInstance()

	httpAddr := ":" + strconv.Itoa(conf.Port)
	e.Logger.Info("API Url:", httpAddr)

	dbService := initDB.Init()
	db, err := dbService.DbConnect("postgres", conf.DbConfig)
	if err != nil {
		e.Logger.Error(err)
		os.Exit(1)
	}
	dm := database.NewDataManager(db)
	defer dm.CloseConnection()

	e.Logger.SetLevel(echoLog.DEBUG)
	InitServices(e, dm, conf)
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
			AllowOrigins:     []string{apiURL, "https://aleksandrkhoroshenin.grafana.net/"},
			AllowCredentials: true,
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
			ExposeHeaders:    []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
		}),
		checkCSRF,
		AccessHitsMiddleware)

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

func InitServices(e *echo.Echo, db database.IDataManager, conf *config.Config) {
	sessionService := security.NewSessionManager(conf.RedisConfig, conf.RedisPass, conf.RedisDbNumber)
	fl := fileLoader.CreateFileLoaderInstance(conf.StorageRegion, conf.StorageEndpoint, conf.StorageBucket)
	userService := user.CreateInstance(db)
	photoService := photo.CreateInstance(db, conf, fl)
	boardService := board.CreateInstance(db)
	boardUsersService := boardUsers.CreateInstance(db)
	cardUsersService := сardUsers.CreateInstance(db)
	cardService := card.CreateInstance(db, fl)
	cardGroupService := cardGroup.CreateInstance(db)
	commentService := comment.CreateInstance(db)
	tagService := tag.CreateInstance(db)
	cardTagsService := cardTags.CreateInstance(db)
	securityService := security.CreateInstance(sessionService)
	e.Use(securityService.CheckSession)
	user.NewUserHandler(e, userService, boardUsersService, cardUsersService, securityService)
	photo.NewPhotoHandler(e, photoService, userService, securityService)
	board.NewBoardHandler(e, userService, boardService, boardUsersService, cardService, cardUsersService, cardGroupService, tagService, cardTagsService, commentService, securityService)
	card.NewCardHandler(e, userService, cardService, cardUsersService, tagService, cardTagsService, commentService, securityService)
	cardGroup.NewCardGroupHandler(e, cardGroupService, securityService)
	comment.NewCommentHandler(e, userService, commentService, securityService)
	tag.NewTagHandler(e, userService, tagService, cardTagsService, securityService)
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

		log.Info("Finish with status code",
			"status_code", ctx.Response().Status,
			"work_time", time.Since(start))
		return h(ctx)
	}
}
