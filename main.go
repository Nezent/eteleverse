package main

import (
	"log"
	"net/http"

	"github.com/Nezent/eteleverse/internal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/api/v1/health", internal.HealthHandler)
	r.Get("/api/v1/hello", internal.HelloHandler)

	port := "8080"
	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
