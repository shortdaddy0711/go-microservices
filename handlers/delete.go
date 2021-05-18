package handlers

import (
	"net/http"

	"github.com/shortdaddy0711/go-microservices/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
// 
// responses:
// 	204: noContent
// 	404: errorResponse
// 	501: errorResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("Handle DELETE Product id:", id)
	w.Header().Set("Content-Type", "application/json")

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found")
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err != nil {
		p.l.Println("[Error] fail to delete product id:", id)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	p.l.Printf("[DEBUG] Deleted product id: %#v\n", id)

	w.WriteHeader(http.StatusNoContent)

}
