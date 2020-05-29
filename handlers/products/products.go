package products

import (
	"log"
)

// Products is a http handler
type Products struct {
	l *log.Logger
}

type keyProduct struct{}

// NewProduct creates new product handler
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}
