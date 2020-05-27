package handlers

import (
	"context"
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

type keyProduct struct{}

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
	prod := r.Context().Value(keyProduct{}).(*data.Product)
	data.AddProducts(prod)
	p.l.Printf("Added Prod: %#v", prod)
}

// UpdateProduct updates the product
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	util.CheckErr(w, err, "Unable to parse id string to number", http.StatusBadRequest)

	prod := r.Context().Value(keyProduct{}).(*data.Product)

	err = prod.FromJSON(r.Body)
	util.CheckErr(w, err, "Unable to unmarshal json", http.StatusBadRequest)

	err = data.UpdateProduct(id, prod)
	util.CheckErr(w, err, "Product not found", http.StatusNotFound)
	p.l.Printf("Update Prod: %#v", prod)
}

// MiddlewareProductAuthentication authenticates product and then proceed to next handler
func (p *Products) MiddlewareProductAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		util.CheckErr(w, err, "Unable to unmarshal JSON", http.StatusBadRequest)

		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
