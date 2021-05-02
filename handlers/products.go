package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shortdaddy0711/go-microservices/data"
)

type Products struct {
	l *log.Logger
}

type KeyProduct struct{}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

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

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

	p.l.Printf("[DEBUG] Inserted product: %#v\n", prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("[ERROR] unable to find product id in URL", r.URL.Path, err)
		http.Error(w, "Missing product id, url should be formatted /products/[id] for put request", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	prod.ID = id

	err = data.UpdateProduct(&prod)
	if err == errors.New("product not found") {
		p.l.Println("[ERROR] product not found", err)
		http.Error(w, "product not found in database", http.StatusNotFound)
		return
	}

	p.l.Printf("[DEBUG] Updated product: %#v\n", prod)

	w.WriteHeader(http.StatusNoContent)
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		if err := prod.FromJSON(r.Body); err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error deserializing product", http.StatusBadRequest)
			return
		}

		if err := prod.Validate(); err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s\n", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
