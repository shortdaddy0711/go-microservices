package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/shortdaddy0711/go-microservices/data"
	"github.com/shortdaddy0711/go-microservices/handlers"

	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "localhost")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()

	// create the handler
	ph := handlers.NewProducts(l, v)

	// create a new serve mux and register the handlers
	r := mux.NewRouter()

	// handlers for API
	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", ph.GetProducts)
	getR.HandleFunc("/products/{id:[0-9]+}", ph.GetAProduct)

	postR := r.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/products", ph.AddProduct)
	postR.Use(ph.MiddlewareProductValidation)

	putR := r.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	putR.Use(ph.MiddlewareProductValidation)

	deleteR := r.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	srv := &http.Server{
		Addr:         *bindAddress,
		Handler:      r,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090")

		if err := srv.ListenAndServe(); err != nil {
			l.Fatalf("Error starting server: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := <-c
	log.Println("Got signal: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	l.Println("Received terminate, graceful shutdown")
	os.Exit(0)
}
