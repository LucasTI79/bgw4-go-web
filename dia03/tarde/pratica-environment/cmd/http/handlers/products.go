package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-put-patch-delete/config"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/apperrors"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/web"
)

type ProductsHandler struct {
	storage domain.Respository
}

func NewProductsHandler(repository domain.Respository) *ProductsHandler {
	return &ProductsHandler{
		storage: repository,
	}
}

func (p *ProductsHandler) Show() http.HandlerFunc {
	// request
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			web.ResponseJSON(w, code, body)
			return
		}

		pr, err := p.storage.GetByID(id)
		if err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": "product not found", "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			web.ResponseJSON(w, code, body)
			return
		}

		code := http.StatusOK
		body := pr

		web.ResponseJSON(w, code, body)
	}
}

func (p *ProductsHandler) UpdateOrCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if token := r.Header.Get("token"); token != config.Config.ApiToken {
			code := http.StatusUnauthorized
			message := map[string]any{"message": "unauthorized", "data": nil}

			web.ResponseJSON(w, code, message)
			return
		}

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			message := map[string]any{"message": "invalid id", "data": nil}

			web.ResponseJSON(w, code, message)
			return
		}

		var reqBody domain.RequestBodyUpdateOrCreate
		if err := web.RequestJSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid request body", "data": nil}

			web.ResponseJSON(w, code, body)
			return
		}

		pr := domain.Product{
			Id:       id,
			Name:     reqBody.Name,
			Type:     reqBody.Type,
			Quantity: reqBody.Quantity,
			Price:    reqBody.Price,
		}
		if err := p.storage.UpdateOrCreate(&pr); err != nil {
			code := http.StatusInternalServerError
			body := map[string]any{"message": "internal server error", "data": nil}

			web.ResponseJSON(w, code, body)
			return
		}

		code := http.StatusOK
		body := pr

		web.ResponseJSON(w, code, body)
	}

}

func (p *ProductsHandler) Update() http.HandlerFunc {
	// request
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			web.ResponseJSON(w, code, body)
			return
		}

		pr, err := p.storage.GetByID(id)
		if err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": "product not found", "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			web.ResponseJSON(w, code, body)
			return
		}

		reqBody := domain.RequestBodyUpdate{
			Name:     pr.Name,
			Type:     pr.Type,
			Quantity: pr.Quantity,
			Price:    pr.Price,
		}
		if err := web.RequestJSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid request body", "data": nil}

			web.ResponseJSON(w, code, body)
			return
		}

		pr = &domain.Product{Id: id, Name: reqBody.Name, Type: pr.Type, Quantity: pr.Quantity, Price: reqBody.Price}
		// -> update
		err = p.storage.Update(id, pr)
		if err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": "product not found", "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			web.ResponseJSON(w, code, body)
			return
		}

		code := http.StatusOK
		body := pr

		web.ResponseJSON(w, code, body)
	}
}

func (p *ProductsHandler) Delete() http.HandlerFunc {
	// request
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := map[string]any{"message": "invalid id", "data": nil}

			web.ResponseJSON(w, code, body)
			return
		}

		err = p.storage.Delete(id)
		if err != nil {
			var code int
			var body map[string]any
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				body = map[string]any{"message": "product not found", "data": nil}
			default:
				code = http.StatusInternalServerError
				body = map[string]any{"message": "internal server error", "data": nil}
			}
			web.ResponseJSON(w, code, body)
			return
		}

		code := http.StatusNoContent
		web.ResponseJSON(w, code, nil)
	}
}
