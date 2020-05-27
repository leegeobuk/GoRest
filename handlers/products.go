package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
		regex := regexp.MustCompile(`/([0-9]+)`)
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid URI, more than one ID", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(w, "Invalid URI, more than one capture group", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "Invalid URI, unable to convert to number", http.StatusBadRequest)
			return
		}

		p.Update(id, w, r)
		return
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

// Update updaates the product
func (p *Products) Update(id int, w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{ID: id}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
}
