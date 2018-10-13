package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/www")))

	srv := &http.Server{
		Addr: ":80",
	}

	log.Printf("Starting to listen at %s", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
