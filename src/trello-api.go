package main

import (
	"2019_2_Shtoby_shto/src/route"
	"flag"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

const (
	port          = ":3001"
	postgreConfig = "postgres://postgres:Aebnm@postgres:5432/db_1?sslmode=disable"
)

var initFlag = flag.Bool("initial start", false, "Check your service")
var httpAddr = flag.String("port", "localhost:8080", "HTTP listen address")

func main() {
	flag.Parse()

	dbInfo := postgreConfig
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()
	initService(db)

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

func initService(db *gorm.DB) {
}
