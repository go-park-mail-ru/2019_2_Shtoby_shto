package main

import (
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/route"
	"flag"
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

	dm := database.DataManager{}
	if err := dm.Init("postgres", postgreConfig); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if *initFlag {
		return
	}

	router := route.NewRouterService()

	s := &http.Server{
		Addr:           *httpAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

func initService() {
}
