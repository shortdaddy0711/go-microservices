package handlers

import (
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
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
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
