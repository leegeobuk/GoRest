package products

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leegeobuk/GoRest/data"
	"github.com/leegeobuk/GoRest/util"
)

// swagger:route GET /products products listAll
// Returns a list of products
// responses:
// 200: productsResponse

// ListAll handles GET request and retrieves all products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	// fetch products from data store
	pl := data.GetProducts()

	// serialize the list to JSON
	err := util.ToJSON(pl, w)
	if util.CheckErr(w, err, "Unable to marshal json", http.StatusInternalServerError) {
		return
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the db
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	prod, err := data.GetProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	err = util.ToJSON(prod, w)

	if util.CheckErr(w, err, "Unable to marshal json", http.StatusInternalServerError) {
		return
	}
}
