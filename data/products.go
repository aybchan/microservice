package data

import "time"

type Product struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
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
