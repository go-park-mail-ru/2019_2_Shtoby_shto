package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"
)

var initFlag = flag.Bool("initial start", false, "Check your service")
var httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")

func main() {
	flag.Parse()

	// TODO:: add timeout for docker
	dbInfo := "postgres://postgres:Aebnm@postgres:5432/db_1?sslmode=disable"
	db, err := sql.Open("postgres", dbInfo)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	if *initFlag {
		return
	}

	s := &http.Server{
		Addr:           *httpAddr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
