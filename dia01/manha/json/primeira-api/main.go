package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
	200-299 -> status de sucesso
	    // 200 - OK
		// 201 - Created
		// 202 - Accept
		// 204 - No Content
	300-399 -> redirecionamentos/cache
	400-499 -> erros do cliente
		// 400 - Bad Request
		// 401 - Unauthorized
		// 403 - Forbidden
		// 404 - Not Found
		// 409 - Conflict
		// 422 - Unprocessable entity
		// 429 - To Many Requests
	500-599 -> erros dos servidor
		// 500 - Internal Server Error
		// 502 - Bad Gateway
*/

func main() {
	router := chi.NewRouter()

	router.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Hello World!"}`))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("erro ao iniciar a api")
	}
}
