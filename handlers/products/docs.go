// Package products classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package products

import "github.com/leegeobuk/GoRest/data"

// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Responses

// A list of products retrieved
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body data.Products
}

// A single product retrieved
// swagger:response productResponse
type productResponse struct {
	// Newly created product
	// in: body
	Body data.Product
}

// swagger:response noContent
type noContent struct{}

// Path variables

// Path variable representing id of product
// swagger:parameters deleteProduct
type productIDParameter struct {
	// id of the product to delete from the db
	// in: path
	// required: true
	ID int `jason:"id"`
}

// Request bodies

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// Errors

// Generic error message
// swagger:response errorResponse
type errorResponse struct {
	// Description of the error
	// in: body
	Body string `json:"message"`
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidation struct {
	// Collection of the errors
	// in: body
	Body []string `json:"messages"`
}
