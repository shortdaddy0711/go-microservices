// Package classification Product API
//
// Documentation for Product API
//
// 	Schemes: http
// 	BasePath: /
// 	Version: 0.0.1
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
//
// swagger:meta
package handlers

import "github.com/shortdaddy0711/go-microservices/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the errorResponseWrapper
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// Data structure representing a single product
// swagger:response productResponse
type productsResponseWrapper struct {
	// Newly created productsResponseWrapper
	// in: body
	Body []data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to update or create
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters updateProduct
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}