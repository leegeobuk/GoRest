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
	}

	// handle an update

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// Get retrieves all products
func (p *Products) Get(w http.ResponseWriter, r *http.Request) {
	pl := data.GetProducts()
	err := pl.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
