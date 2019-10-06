package main

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/photo"
	"2019_2_Shtoby_shto/src/dicts/user"
	transport "2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/route"
	"2019_2_Shtoby_shto/src/security"
	"2019_2_Shtoby_shto/src/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"strconv"
)

const (
	deployEnvVar = "DEPLOYAPI"
)

var (
	initFlag = flag.Bool("initial start", false, "Check your service")
	httpAddr = flag.String("address", ":8080", "HTTP listen address")
)

var (
	transportService transport.Handler
	securityService  security.Security
	userService      user.HandlerUserService
	photoService     photo.HandlerPhotoService
	dbService        database.InitDBManager
)

var logger *log.Logger

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

	initService(dm, conf)
	srv := newServer(logger)
	if *initFlag {
		return
	}

	//TODO::great shutdown
	switch conf.Port {
	case 443:
		if err := srv.ListenAndServeTLS("keys/server.crt", "keys/server.key");
		   err != http.ErrServerClosed {
			logger.Fatalf("HTTPS server ListenAndServe: %v", err)
		}
	default:
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}
}

func newServer(logger *log.Logger) *http.Server {
	router := AccessLogMiddleware(AccessCORS(route.NewRouterService(securityService)))

	logger.Println("serving on", *httpAddr)

	return &http.Server{
		Addr:           *httpAddr,
		Handler:        router,
		ErrorLog:       logger,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func initService(db database.IDataManager, conf *config.Config) {
	sessionService := security.NewSessionManager(conf.RedisConfig, conf.RedisPass, conf.RedisDbNumber)
	userService = user.CreateInstance(db)
	photoService = photo.CreateInstance(db)
	transportService = transport.CreateInstance(sessionService)
	securityService = security.CreateInstance(sessionService, userService, photoService)
}

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessLogMiddleware", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func AccessCORS(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.SetHeaders(&w)
		if (*r).Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
