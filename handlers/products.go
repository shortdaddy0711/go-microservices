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

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shortdaddy0711/go-microservices/data"
)

// Products handler for get/update products
type Products struct {
	l *log.Logger
	v *data.Validation
}

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// NewProducts returns a new products handler with the given logger and validator
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// getProductID returns the product ID from request path
func getProductID(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}

var ErrInvalidProductPath = fmt.Errorf("invalid path, path should be /products/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		v := data.NewValidation()

		if err := data.FromJSON(prod, r.Body); err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error deserializing product", http.StatusBadRequest)
			return
		}

		if err := v.Validate(prod); err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s\n", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
