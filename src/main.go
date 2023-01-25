package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	server := &http.Server{
		Addr:           ":8080",
		Handler:        limit(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())

}
