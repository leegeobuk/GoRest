package products

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leegeobuk/GoRest/data"
	"github.com/leegeobuk/GoRest/util"
)

// swagger:route DELETE /products{id} products deleteProduct
// Deletes a product
// responses:
// 201: noContent

// Delete handles DELETE requests and removes item from the database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if util.CheckErr(w, err, "Unable to parse id string to number", http.StatusBadRequest) {
		return
	}

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if util.CheckErr(w, err, "Internal server error", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
