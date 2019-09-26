package main

import (
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/user"
	transport "2019_2_Shtoby_shto/src/handle"
	"2019_2_Shtoby_shto/src/route"
	"2019_2_Shtoby_shto/src/security"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	port = ":3001"
	//postgreConfig = "postgres:Aebnm@postgres:5432/db_1?sslmode=disable"
	//postgreConfig = "postgres://postgres:Aebnm@postgres:5432/db_1?sslmode=disable"
	postgreConfig = "host='localhost' port=5432 user=postgres dbname='trello' sslmode=disable password='1111'"
)

var initFlag = flag.Bool("initial start", false, "Check your service")
var httpAddr = flag.String("address", "localhost:8080", "HTTP listen address")

var (
	transportService transport.Handler
	securityService  security.Security
	userService      user.UserHandler
	dbService        database.IDataManager
)

func main() {
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	// TODO::add context with dm and sessionId
	dm := &database.DataManager{}
	dbService = database.Init()
	if err := dbService.DbConnect(dm, "postgres", postgreConfig); err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	initService(dm)
	srv := newServer(logger)
	if *initFlag {
		return
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func newServer(logger *log.Logger) *http.Server {
	router := AccessLogMiddleware(route.NewRouterService(securityService))
	return &http.Server{
		Addr:           *httpAddr,
		Handler:        router,
		ErrorLog:       logger,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func initService(db *database.DataManager) {
	sessionService := security.NewSessionManager("localhost:6379", "", 0)
	userService = user.CreateInstance(db)
	transportService = transport.CreateInstance(sessionService)
	securityService = security.CreateInstance(sessionService, userService)
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
