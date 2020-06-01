package handlers

import (
	"log"

	"github.com/leegeobuk/GoServer/GoServer/handlers/products"
)

// NewProduct creates new product handler
func NewProduct(l *log.Logger) *products.Products {
	return products.NewProduct(l)
}
