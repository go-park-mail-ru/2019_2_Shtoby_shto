package main

import (
	"2019_2_Shtoby_shto/src/database"
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
	port          = ":3001"
	postgreConfig = "postgres://postgres:Aebnm@postgres:5432/db_1?sslmode=disable"
)

var initFlag = flag.Bool("initial start", false, "Check your service")
var httpAddr = flag.String("address", "localhost:8080", "HTTP listen address")

var (
	transportService transport.Handler
	securityService  security.Security
)

func main() {
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	initService()
	srv := newServer(logger)
	if *initFlag {
		return
	}

	dm := database.DataManager{}
	if err := dm.NewDataManager("postgres", postgreConfig); err != nil {
		logger.Fatal(err)
		os.Exit(1)
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

func initService() {
	sessionService := security.NewSessionManager("localhost:6379", "", 0)

	transportService = transport.CreateInstance(sessionService)
	securityService = security.CreateInstance(sessionService)
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
