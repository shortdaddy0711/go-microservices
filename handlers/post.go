package handlers

import (
	"net/http"

	"github.com/shortdaddy0711/go-microservices/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
// 	200: productResponse
// 	422: errorVaidation
// 	501: errorResponse

// Create handles POST requests to add new product
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

	p.l.Printf("[DEBUG] Inserted product: %#v\n", prod)
}
