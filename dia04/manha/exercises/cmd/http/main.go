package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-pratica-web/cmd/http/routes"
	"github.com/lucasti79/bgw4-pratica-web/config"
)

func main() {
	config.Init()

	r := chi.NewRouter()

	routes := routes.NewRoutes(r)
	routes.MapRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
