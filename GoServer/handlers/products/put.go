package products

import (
	"net/http"

	"github.com/leegeobuk/GoServer/GoServer/data"
	"github.com/leegeobuk/GoServer/GoServer/util"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContent
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update a product
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(keyProduct{}).(*data.Product)

	err := data.UpdateProduct(prod)
	if util.CheckErr(w, err, "Product not found", http.StatusNotFound) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
	p.l.Printf("Update Prod: %#v", prod)
}
