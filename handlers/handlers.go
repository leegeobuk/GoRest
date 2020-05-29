package handlers

import (
	"log"

	"github.com/leegeobuk/GoRest/handlers/products"
)

// NewProduct creates new product handler
func NewProduct(l *log.Logger) *products.Products {
	return products.NewProduct(l)
}
