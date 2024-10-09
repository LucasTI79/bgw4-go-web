package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-pratica-web/cmd/http/handlers"
	"github.com/lucasti79/bgw4-pratica-web/cmd/http/middlewares"
	"github.com/lucasti79/bgw4-pratica-web/internal/domain"
	"github.com/lucasti79/bgw4-pratica-web/internal/products"
)

func (r *routes) buildProductRoutes() {
	r.router.Route("/products", func(r chi.Router) {
		r.Use(middlewares.Auth)
		storage := make(map[int]*domain.Product)
		repo := products.NewRepository(storage)
		service := products.NewService(repo)
		productHandler := handlers.NewProductHandler(service)

		r.Get("/", productHandler.Index)
		r.Get("/{productId}", productHandler.Show)
		r.Post("/", productHandler.Create)
		r.Put("/{productId}", productHandler.Update)
		r.Delete("/{productId}", productHandler.Delete)
	})
}
