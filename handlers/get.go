package handlers

import (
	"net/http"

	"github.com/shortdaddy0711/go-microservices/data"
)

// swagger:route GET /products products listProduts
// Returns a list of products
// responses:
// 	200: productsResponse
// 	404: errorResponse

// GetProducts returns the products from the DB
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	w.Header().Set("Content-Type", "application/json")

	pl := data.GetProducts()
	err := data.ToJSON(pl, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Returns a product of the given id
// Responses:
// 	200: productResponse
// 	404: errorResponse
//
// GetAProduct returns a product from the DB
func (p *Products) GetAProduct(w http.ResponseWriter, r *http.Request) {

	id := getProductID(r)
	p.l.Println("Handle GET Product id: ", id)
	w.Header().Set("Content-Type", "application/json")

	prod, err := data.GetProductByID(id)

	if err != nil {
		if err == data.ErrProductNotFound {
			p.l.Println("[ERROR] no match", err)

			w.WriteHeader(http.StatusNotFound)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		} else {
			p.l.Println("[ERROR] fetching product", err)

			w.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
