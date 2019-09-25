package main

import (
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/route"
	"2019_2_Shtoby_shto/src/security"
	"flag"
	"github.com/go-redis/redis"
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

func main() {
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	sessionManager := security.NewSessionManager()
	println(sessionManager.Check(&security.SessionID{ID: "123"}))
	srv := newServer(logger)
	if *initFlag {
		return
	}

	dm := database.DataManager{}
	if err := dm.Init("postgres", postgreConfig); err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	//idleConnsClosed := make(chan struct{})
	//go func() {
	//	sigint := make(chan os.Signal, 1)
	//	signal.Notify(sigint, os.Interrupt)
	//	<-sigint
	//	if err := srv.Shutdown(context.Background()); err != nil {
	//		logger.Printf("HTTP server Shutdown: %v", err)
	//	}
	//	close(idleConnsClosed)
	//}()
	//if err := srv.ListenAndServe(); err != http.ErrServerClosed {
	//	logger.Fatalf("HTTP server ListenAndServe: %v", err)
	//}
	//<-idleConnsClosed
}

func newServer(logger *log.Logger) *http.Server {
	router := route.NewRouterService()
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
}
