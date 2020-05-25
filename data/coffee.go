package data

import (
	"encoding/json"
	"io"
	"time"
)

// Coffee defines the structure for an API product
type Coffee struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	Created     string  `json:"-"`
	Updated     string  `json:"-"`
	Deleted     string  `json:"-"`
}

// Products holds multiple product
type Products []*Coffee

// ToJSON parses products to JSON format
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns slice of products
func GetProducts() Products {
	return productList
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		Created:     time.Now().UTC().String(),
		Updated:     time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		Created:     time.Now().UTC().String(),
		Updated:     time.Now().UTC().String(),
	},
}
