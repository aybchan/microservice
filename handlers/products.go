package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		p.putProduct(rw, r)
        return
	}

	// handle unimplemented verbs
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) putProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling product PUT")
	path := r.URL.Path

	rxp := `/([0-9]+)`
	pattern := regexp.MustCompile(rxp)
	matches := pattern.FindAllStringSubmatch(path, -1)

	if len(matches) != 1 {
		http.Error(rw, "Invalid URI for PUT method", http.StatusBadRequest)
		return
	}

	if len(matches[0]) != 2 {
		http.Error(rw, "Invalid URI for PUT method", http.StatusBadRequest)
		return
	}

	match := matches[0][1]
	id, err := strconv.Atoi(match)
	if err != nil {
		http.Error(rw, "Could not convert path to int", http.StatusBadRequest)
	}

	p.l.Println(path, id)
    err = p.updateProduct(id, rw, r)
    if err == data.ProductNotFound {
        http.Error(rw, "Product not found", http.StatusBadRequest)
        return
    }
    if err != nil {
        http.Error(rw, "Could not return", http.StatusInternalServerError)

    }
	return
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
