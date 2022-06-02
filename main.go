package main

import (
	"log"
	"net/http"
	"os"

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

	http.ListenAndServe(":9090", sm)
}
