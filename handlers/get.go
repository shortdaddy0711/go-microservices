package handlers

import (
	"net/http"

	"github.com/shortdaddy0711/go-microservices/data"
)

// swagger:route GET /products products listProduts
// Returns a list of products
// responses:
// 	200: productsResponseWrapper

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	w.Header().Set("Content-Type", "application/json")

	pl := data.GetProducts()
	err := pl.ToJSON(w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
		http.Error(w, "Error serializing product", http.StatusInternalServerError)
	}
}
