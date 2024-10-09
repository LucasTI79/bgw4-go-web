package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-go-web-parameters/cmd/http/handlers"
	"github.com/lucasti79/bgw4-go-web-parameters/internal/domain"
)

func main() {
	r := chi.NewRouter()

	employeeHandler := handlers.HandlerEmployee{}
	db := make(map[int]*domain.Product)
	productsHandler := handlers.NewProductsHandler(db)

	r.Get("/api/employees/{id}", employeeHandler.GetById())
	r.Route("/api/products", func(pg chi.Router) {
		pg.Post("/", productsHandler.Create())
	})

	// INDEX -> buscar todos os produtos
	// "/api/products"
	// SHOW -> buscar um unico produto
	// "/api/products/{productId}"
	// STORE -> criar um produto
	// "/api/products"
	// UPDATE -> atualizar um produto
	// "/api/products/{productId}"
	// DELETE -> deletar um produto
	// "/api/products/{productId}"

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
