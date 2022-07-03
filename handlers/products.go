package handlers

import (
	"log"
	"net/http"
	"strconv"

    "github.com/gorilla/mux"
	"github.com/aybchan/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

	p.l.Println("Handling product PUT, updating product ID", id)

	idint, err := strconv.Atoi(id)

	if err != nil {
		http.Error(rw, "Could not convert path to int", http.StatusBadRequest)
	}

	p.updateProduct(idint, rw, r)

	return
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling GET product")
	productList := data.GetProducts()

	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Encoding failed", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST product")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Decoding failed", http.StatusBadRequest)
	}
	data.AddProduct(product)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) error {
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Decoding failed", http.StatusBadRequest)
        return err
	}

	err = data.UpdateProduct(id, product)
	if err != nil {
		http.Error(rw, "Product not found with given ID", http.StatusBadRequest)
        return err
	}

    return nil
}
