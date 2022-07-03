package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

    "github.com/gorilla/mux"
	"github.com/aybchan/microservice/handlers"
)

func main() {
	// create handlers, logger for dependency injection
	l := log.New(os.Stdout, "api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	byeHandler := handlers.NewBye(l)
	productHandler := handlers.NewProducts(l)

	// register handlers to servemux
	//sm := http.NewServeMux()
    sm := mux.NewRouter()
	sm.Handle("/hello", helloHandler)
	sm.Handle("/bye", byeHandler)

    getRouter := sm.Methods(http.MethodGet).Subrouter()
    getRouter.HandleFunc("/", productHandler.GetProducts)

    putRouter := sm.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)

    postRouter := sm.Methods(http.MethodPost).Subrouter()
    postRouter.HandleFunc("/", productHandler.AddProduct)

	// manually create http server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  500 * time.Millisecond,
		ReadTimeout:  10 * time.Millisecond,
		WriteTimeout: 10 * time.Millisecond,
	}

	// non-blocking serve
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// create channel for shutdown
	sigChan := make(chan os.Signal)

	// send message into channel on kill signals
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// blocking wait for signal in channel
	sig := <-sigChan
	l.Println("interrupt or kill signal received in chan, gracefully shutting down", sig)

	//  shutdown
	tc, _ := context.WithTimeout(context.Background(), time.Second*300)
	s.Shutdown(tc)
}
