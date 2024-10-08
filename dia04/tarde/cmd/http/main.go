package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-put-patch-delete/cmd/http/handlers"
	"github.com/lucasti79/bgw4-put-patch-delete/cmd/http/middlewares"
	"github.com/lucasti79/bgw4-put-patch-delete/config"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/storage"
)

// TDD -> Teste depois do deploy

func main() {
	config.InitEnvironment()
	r := chi.NewRouter()

	db := storage.NewProductsStorage(map[int]storage.ProductAttributes{})
	productsHandler := handlers.NewProductsHandler(db)
	healthHandler := handlers.NewHealthHandler()

	r.Use(middlewares.CheckTime)
	r.Get("/health", healthHandler.Health())
	// middleware global

	r.Route("/api/products", func(pg chi.Router) {
		// middleware global apenas para rotas nesse grupo
		// pg.Use(middlewares.CheckTime)
		// middleware para rota em expecifico
		// pg.With(middlewares.CheckTime).Get("/", productsHandler.Get())
		pg.Use(middlewares.Auth)
		pg.Get("/", productsHandler.Get())
		pg.Get("/{id}", productsHandler.Show())
		pg.Put("/{id}", productsHandler.UpdateOrCreate())
		pg.Patch("/{id}", productsHandler.Update())
		pg.Delete("/{id}", productsHandler.Delete())
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
