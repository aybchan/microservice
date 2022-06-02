package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aybchan/microservice/handlers"
)

func main() {
	// create handlers
	l := log.New(os.Stdout, "api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	byeHandler := handlers.NewBye(l)

	// register handlers to servemux
	sm := http.NewServeMux()
	sm.Handle("/hello", helloHandler)
	sm.Handle("/bye", byeHandler)

	// manually create http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  500 * time.Millisecond,
		ReadTimeout:  10 * time.Millisecond,
		WriteTimeout: 10 * time.Millisecond,
	}

	s.ListenAndServe()
}
