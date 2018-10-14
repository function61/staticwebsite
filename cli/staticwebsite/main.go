package main

import (
	"github.com/function61/gokit/ossignal"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/www")))

	srv := &http.Server{
		Addr: ":80",
	}

	go func() {
		log.Printf("Starting to listen at %s", srv.Addr)

		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ListenAndServe() error: %s", err.Error())
		}
	}()

	log.Printf("Got %s; shutdown requested", ossignal.WaitForInterruptOrTerminate())

	if err := srv.Shutdown(nil); err != nil {
		log.Fatal(err)
	}
}
