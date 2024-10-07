package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-go-web-parameters/cmd/server/handlers"
)

func main() {
	r := chi.NewRouter()

	employeeHandler := handlers.HandlerEmployee{}

	r.Get("/api/employees/{id}", employeeHandler.GetById())

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
