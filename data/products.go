package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
}

type Products []*Product

func (ps Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(ps)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	var i int
	for i = 0; i < len(productList); i++ {
		if productList[i].ID == id {
			p.ID = id
			productList[i] = p
			return nil
		}
	}
	return ProductError{Message: "Product with ID not found"}
}

type ProductError struct {
	Message string
}

func (pe ProductError) Error() string {
	return fmt.Sprintf(pe.Message)
}

func getNextID() int {
	return productList[len(productList)-1].ID + 1
}

func GetProducts() Products {
	return productList
}

var productList = Products{
	{
		ID:          0,
		Name:        "Espresso",
		Description: "Short coffee",
		Price:       1.00,
		SKU:         "coffee00",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          1,
		Name:        "Latte",
		Description: "Foamy milk and espresso",
		Price:       2.45,
		SKU:         "coffee01",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
