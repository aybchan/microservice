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
	productList := data.GetProducts()
	err := productList.ToJSON(rw)

	// using encoder over marshal avoids placement in memory
	//j, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "Encoding failed", http.StatusInternalServerError)
	}
	//rw.Write(j)
}
