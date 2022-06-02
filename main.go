package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(http.ResponseWriter, *http.Request) {
		log.Println("hello world")
	})
	http.ListenAndServe(":9090", nil)
}
