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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// handle unimplemented verbs
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling GET product")
	productList := data.GetProducts()

	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Encoding failed", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST product")

	product := &data.Product{}
	p.l.Println(r.Body)
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Decoding failed", http.StatusBadRequest)
	}
	p.l.Printf("Product: %#v\n", product)
}
