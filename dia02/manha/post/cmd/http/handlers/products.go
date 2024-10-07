package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lucasti79/bgw4-go-web-parameters/internal/domain"
)

type ProductsHandler struct {
	storage map[int]*domain.Product
}

func NewProductsHandler(db map[int]*domain.Product) *ProductsHandler {
	return &ProductsHandler{
		storage: db,
	}
}

func (c *ProductsHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "123456" {
			code := http.StatusUnauthorized // 401
			body := &domain.ResponseBodyProduct{Message: "Unauthorized", Data: nil, Error: true}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		var reqBody domain.RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			code := http.StatusUnprocessableEntity
			body := &domain.ResponseBodyProduct{
				Message: "Unprocessable entity",
				Data:    nil,
				Error:   true,
			}

			// precisa ser definido antes do WriteHeader, senao nao funfa
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		pr := &domain.Product{
			Id:       len(c.storage) + 1,
			Name:     reqBody.Name,
			Type:     reqBody.Type,
			Quantity: reqBody.Quantity,
			Price:    reqBody.Price,
		}
		// -> save product
		c.storage[pr.Id] = pr

		code := http.StatusCreated
		body := &domain.ResponseBodyProduct{
			Message: "Product created",
			Data: &domain.Product{
				Id:       pr.Id,
				Name:     pr.Name,
				Type:     pr.Type,
				Quantity: pr.Quantity,
				Price:    pr.Price,
			},
			Error: false,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
	}
}
