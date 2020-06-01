package products

import (
	"net/http"

	"github.com/leegeobuk/GoRest/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// AddProducts adds product to products list
func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(keyProduct{}).(*data.Product)
	data.AddProduct(prod)
	p.l.Printf("Added Prod: %#v", prod)
}
