package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aybchan/microservice/handlers"
)

func main() {
	// create handler
	l := log.New(os.Stdout, "api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)

	// register handler to servemux
	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)

	//http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
	//	log.Println("good bye")

	//	os.Exit(3)
	//})

	http.ListenAndServe(":9090", sm)
}
