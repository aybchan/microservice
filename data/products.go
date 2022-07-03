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
    idx, err := findProduct(id)
    if err != nil {
        return err
    }

    p.ID = id
    productList[idx] = p
    return nil

}

func findProduct(id int) (int, error) {
    for i, product := range(productList) {
		if product.ID == id {
            return i, nil
        }
    }
    return -1, ProductNotFound
}

type ProductError struct {
	Message string
}
var ProductNotFound = ProductError{Message: "Product with ID not found"}

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
