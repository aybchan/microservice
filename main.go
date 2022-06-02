package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello world")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "something went wrong", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "hello %s", d)
	})

	http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("good bye")

		os.Exit(3)
	})

	http.ListenAndServe(":9090", nil)
}
