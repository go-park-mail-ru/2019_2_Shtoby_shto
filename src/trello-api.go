package main

import (
	"2019_2_Shtoby_shto/src/database"
	"flag"
	"log"
	"net/http"
	"time"
)

var initFlag = flag.Bool("initial start", false, "Check your service")
var httpAddr = flag.String("port", ":8080", "HTTP listen address")

func main() {
	flag.Parse()

	//db, err := database.NewService()

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
