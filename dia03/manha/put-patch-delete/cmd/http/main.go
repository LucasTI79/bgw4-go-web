package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-put-patch-delete/cmd/http/handlers"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/storage"
)

// TDD -> Teste depois do deploy

func main() {
	r := chi.NewRouter()

	db := storage.NewProductsStorage(map[int]storage.ProductAttributes{})
	productsHandler := handlers.NewProductsHandler(db)

	r.Route("/api/products", func(pg chi.Router) {
		pg.Get("/{id}", productsHandler.Show())
		pg.Put("/{id}", productsHandler.UpdateOrCreate())
		pg.Patch("/{id}", productsHandler.Update())
		pg.Delete("/{id}", productsHandler.Delete())
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
