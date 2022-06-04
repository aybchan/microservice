package handlers

import (
	"log"
	"net/http"

	"github.com/aybchan/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle REST verbs
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle unimplemented verbs
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()

	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Encoding failed", http.StatusInternalServerError)
	}
}
