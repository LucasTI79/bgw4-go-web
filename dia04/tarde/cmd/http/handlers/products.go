package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (p *ProductsHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pr, err := p.storage.Get()

		if err != nil {
			code := http.StatusInternalServerError
			web.ResponseJSON(w, code, nil, "internal server error")
			return
		}

		code := http.StatusOK
		if len(pr) == 0 {
			code = http.StatusNoContent
		}

		web.ResponseJSON(w, code, pr, "products list")
	}
}

func (p *ProductsHandler) Show() http.HandlerFunc {
	// request
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			web.ResponseJSON(w, code, nil, "invalid id")
			return
		}

		pr, err := p.storage.GetByID(id)
		if err != nil {
			var code int
			var message string
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				message = "product not found"
			default:
				code = http.StatusInternalServerError
				message = "internal server error"
			}
			web.ResponseJSON(w, code, nil, message)
			return
		}

		code := http.StatusOK
		web.ResponseJSON(w, code, pr, "product")
	}
}

func (p *ProductsHandler) UpdateOrCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			message := "invalid id"

			web.ResponseJSON(w, code, nil, message)
			return
		}

		var reqBody domain.RequestBodyUpdateOrCreate
		if err := web.RequestJSON(r, &reqBody); err != nil {
			code := http.StatusBadRequest
			message := "invalid request body"

			web.ResponseJSON(w, code, nil, message)
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
			message := "internal server error"

			web.ResponseJSON(w, code, nil, message)
			return
		}

		code := http.StatusOK
		web.ResponseJSON(w, code, pr, "product")
	}

}

func (p *ProductsHandler) Update() http.HandlerFunc {
	// request
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			message := "invalid id"

			web.ResponseJSON(w, code, nil, message)
			return
		}

		pr, err := p.storage.GetByID(id)
		if err != nil {
			var code int
			var message string
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				message = "product not found"
			default:
				code = http.StatusInternalServerError
				message = "internal server error"
			}
			web.ResponseJSON(w, code, nil, message)
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
			message := "invalid request body"

			web.ResponseJSON(w, code, nil, message)
			return
		}

		pr = &domain.Product{Id: id, Name: reqBody.Name, Type: pr.Type, Quantity: pr.Quantity, Price: reqBody.Price}
		// -> update
		err = p.storage.Update(id, pr)
		if err != nil {
			var code int
			var message string
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				message = "product not found"
			default:
				code = http.StatusInternalServerError
				message = "internal server error"
			}
			web.ResponseJSON(w, code, nil, message)
			return
		}

		code := http.StatusOK
		web.ResponseJSON(w, code, pr, "product")
	}
}

func (p *ProductsHandler) Delete() http.HandlerFunc {
	// request
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			message := "invalid id"

			web.ResponseJSON(w, code, nil, message)
			return
		}

		err = p.storage.Delete(id)
		if err != nil {
			var code int
			var message string
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				code = http.StatusNotFound
				message = "product not found"
			default:
				code = http.StatusInternalServerError
				message = "internal server error"
			}
			web.ResponseJSON(w, code, nil, message)
			return
		}

		code := http.StatusNoContent
		web.ResponseJSON(w, code, nil, "product deleted")
	}
}
