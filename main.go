package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/shortdaddy0711/go-microservices/handlers"

	"github.com/nicholasjackson/env"
	"github.com/gorilla/mux"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "localhost")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	r := mux.NewRouter()
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("/", ph.GetProducts).Methods("GET")
	s.HandleFunc("/", ph.AddProduct).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct).Methods(http.MethodPut)

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
			l.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	l.Println("Received terminate, graceful shutdown")
	os.Exit(0)
}
