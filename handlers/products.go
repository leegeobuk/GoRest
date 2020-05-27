package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leegeobuk/GoRest/data"
	"github.com/leegeobuk/GoRest/util"
)

// Products defines product handler
type Products struct {
	l *log.Logger
}

// NewProduct creates new product handler
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts retrieves all products in product list
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	// fetch products from data store
	pl := data.GetProducts()

	// serialize the list to JSON
	err := pl.ToJSON(w)
	util.CheckErr(w, err, "Unable to marshal json", http.StatusInternalServerError)
}

// AddProducts adds product to products list
func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	util.CheckErr(w, err, "Unable to unmarshal json", http.StatusBadRequest)

	data.AddProducts(prod)
	p.l.Printf("Prod: %#v", prod)
}

// UpdateProduct updates the product
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	util.CheckErr(w, err, "Unable to parse id string to number", http.StatusBadRequest)

	prod := &data.Product{ID: id}

	err = prod.FromJSON(r.Body)
	util.CheckErr(w, err, "Unable to unmarshal json", http.StatusBadRequest)

	err = data.UpdateProduct(id, prod)
	util.CheckErr(w, err, "Product not found", http.StatusNotFound)
}
