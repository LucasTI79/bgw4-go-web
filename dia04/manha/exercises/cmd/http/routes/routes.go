package routes

import "github.com/go-chi/chi/v5"

type Routes interface {
	MapRoutes(*chi.Mux)
}

type routes struct {
	router *chi.Mux
}

func NewRoutes(router *chi.Mux) Routes {
	return &routes{router: router}
}

func (r *routes) MapRoutes(router *chi.Mux) {
	r.router = router
	r.buildProductRoutes()
}
