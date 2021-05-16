package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shortdaddy0711/go-microservices/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
// responses:
// 	201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("[ERROR] unable to find product id in URL", r.URL.Path, err)
		http.Error(w, "Missing product id, url should be formatted /products/[id] for delete request", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle DELETE Product", id)
	w.Header().Set("Content-Type", "application/json")

	err = data.DeleteProduct(id)
	if err == errors.New("product not found") {
		p.l.Println("[ERROR] product not found", err)
		http.Error(w, "product not found in database", http.StatusNotFound)
		return
	}

	p.l.Printf("[DEBUG] Deleted product id: %#v\n", id)

	w.WriteHeader(http.StatusNoContent)

}
