package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
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
type Products []*Product

// ErrProductNotFound thrown when product doesn't exist
var ErrProductNotFound = errors.New("Product not found")

// FromJSON parses JSON to Go value
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// ToJSON parses products to JSON format
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns slice of products
func GetProducts() Products {
	return productList
}

// AddProducts adds a product to product list
func AddProducts(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct updates the product with given id
func UpdateProduct(id int, p *Product) error {
	i, err := findIndex(id)

	if err != nil {
		return err
	}

	productList[i] = p
	return nil
}

func findIndex(id int) (int, error) {
	for i, p := range productList {
		if id == p.ID {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
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
