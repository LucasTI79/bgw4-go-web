package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-pratica-web/cmd/http/handlers/dtos"
	"github.com/lucasti79/bgw4-pratica-web/internal/products"
	"github.com/lucasti79/bgw4-pratica-web/pkg/apperrors"
	"github.com/lucasti79/bgw4-pratica-web/pkg/web"
)

type ProductHandler struct {
	ProductService products.Service
}

func NewProductHandler(productService products.Service) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) Index(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(products) == 0 {
		web.SucessResponse(w, http.StatusNoContent, products)
		return
	}

	web.SucessResponse(w, http.StatusOK, products)
	return
}

func (h *ProductHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "productId"))
	if err != nil {
		web.ErrorResponse(w, http.StatusBadRequest, "product id must be a number")
		return
	}

	product, err := h.ProductService.FindByID(id)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			{
				web.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("product with id %d not found", id))
				return
			}
		default:
			{
				web.ErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	web.SucessResponse(w, http.StatusOK, product)
	return
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product dtos.CreateProductInput
	if err := web.RequestJSON(r, &product); err != nil {
		web.ErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// if !product.IsValid() {
	// 	web.ErrorResponse(w, http.StatusBadRequest, "invalid product")
	// 	return
	// }

	newProduct, err := h.ProductService.Create(product.ToDomain())
	if err != nil {
		switch {
		case errors.Is(err, products.ErrCodeAlreadyExists):
			{
				web.ErrorResponse(w, http.StatusConflict, err.Error())
				return
			}
		default:
			{
				web.ErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	web.SucessResponse(w, http.StatusCreated, newProduct)
	return
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	var product dtos.UpdateProductInput

	id, err := strconv.Atoi(chi.URLParam(r, "productId"))
	if err != nil {
		web.ErrorResponse(w, http.StatusBadRequest, "product id must be a number")
		return
	}

	if err := web.RequestJSON(r, &product); err != nil {
		web.ErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// if !product.IsValid() {}

	updatedProduct, err := h.ProductService.Update(id, product.ToDomain())

	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			{
				web.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("product with id %d not found", id))
				return
			}
		case errors.Is(err, products.ErrCodeAlreadyExists):
			{
				web.ErrorResponse(w, http.StatusConflict, err.Error())
				return
			}
		default:
			{
				web.ErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	web.SucessResponse(w, http.StatusOK, updatedProduct)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "productId"))
	if err != nil {
		web.ErrorResponse(w, http.StatusBadRequest, "product id must be a number")
		return
	}

	if err := h.ProductService.Delete(id); err != nil {
		switch {
		case errors.Is(err, apperrors.ErrNotFound):
			{
				web.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("product with id %d not found", id))
				return
			}
		default:
			{
				web.ErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	web.SucessResponse(w, http.StatusNoContent, nil)
	return
}
