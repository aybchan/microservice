package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	http.ListenAndServe(":9090", nil)
}
