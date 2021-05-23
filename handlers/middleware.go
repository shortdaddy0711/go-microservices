package handlers

import (
	"context"
	"net/http"

	"github.com/shortdaddy0711/go-microservices/data"
)

// MiddlewareProductValidation validates the product in the request and calls next if ok
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		if err := data.FromJSON(prod, r.Body); err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		if errs := p.v.Validate(prod); len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// validate the product
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// call the next handle, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
