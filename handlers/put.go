package handlers

import (
	"net/http"

	"github.com/shortdaddy0711/go-microservices/data"
)

// swagger: route PUT /products products UpdateProduct
// Update a products details
//
// responses:
// 	204: noContentResponse
// 	404: errorResponse
// 	422: errorValidation
//
// UpdateProduct update a product
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in the DB"}, w)
		return
	}

	p.l.Printf("[DEBUG] Updated product: %#v\n", prod)
	w.WriteHeader(http.StatusNoContent)
}
