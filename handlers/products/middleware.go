package products

import (
	"context"
	"fmt"
	"net/http"

	"github.com/leegeobuk/GoRest/data"
	"github.com/leegeobuk/GoRest/util"
)

// MiddlewareProductAuthentication validates product in the request and calls next if ok
func (p *Products) MiddlewareProductAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := util.FromJSON(prod, r.Body)
		if util.CheckErr(w, err, "Unable to unmarshal JSON", http.StatusBadRequest) {
			return
		}

		err = prod.Validate()
		if util.CheckErr(w, err, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest) {
			return
		}

		// add the product to context
		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
