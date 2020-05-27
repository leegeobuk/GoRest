package handlers

import (
	"log"
	"net/http"

	"github.com/leegeobuk/GoRestStdlib/data"
)

// Products defines product handler
type Products struct {
	l *log.Logger
}

// NewProduct creates new product handler
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.Get(w, r)
		return
	} else if r.Method == http.MethodPost {
		p.Post(w, r)
		return
	}
	if r.Method == http.MethodPut {
		
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// Get retrieves all products
func (p *Products) Get(w http.ResponseWriter, r *http.Request) {
	// fetch products from data store
	pl := data.GetProducts()

	// serialize the list to JSON
	err := pl.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// Post product to products list
func (p *Products) Post(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProducts(prod)
	p.l.Printf("Prod: %#v", prod)
}
