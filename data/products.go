package data

import "time"

type Product struct {
	ID          int32
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedAt   string
	UpdatedAt   string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
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
