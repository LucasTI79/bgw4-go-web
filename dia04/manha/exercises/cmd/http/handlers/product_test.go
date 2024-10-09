package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-pratica-web/cmd/http/handlers"
	"github.com/lucasti79/bgw4-pratica-web/cmd/http/routes"
	"github.com/lucasti79/bgw4-pratica-web/internal/domain"
	"github.com/lucasti79/bgw4-pratica-web/internal/products"
	"github.com/lucasti79/bgw4-pratica-web/pkg/testutils"
	"github.com/lucasti79/bgw4-pratica-web/pkg/web"
	"github.com/stretchr/testify/assert"
)

func TestProductHandler_Index(t *testing.T) {
	t.Run("should return a list of products and status code OK", func(t *testing.T) {
		// given
		product1 := domain.Product{
			Id:         1,
			Name:       "product 1",
			Quantity:   new(int),
			Code:       "code1",
			Published:  new(bool),
			Expiration: "2024-10-09",
			Price:      0.0,
		}

		db := map[int]*domain.Product{
			1: &product1,
		}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res := httptest.NewRecorder()
		hd.Index(res, req)

		// then

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		// expectedData := []domain.Product{product1}
		// expectedLength := 1
		getResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &getResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		// require.Equal(t, expectedLength, len(getResponse.Data.([]domain.Product)))
		// require.ElementsMatch(t, expectedData, getResponse.Data)
	})
	t.Run("should return no content when there are no products", func(t *testing.T) {
		// given
		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res := httptest.NewRecorder()
		hd.Index(res, req)

		// then

		expectedCode := http.StatusNoContent
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
	})
	t.Run("should return a unauthorized code when a request doesnt have a token", func(t *testing.T) {
		// given
		os.Setenv("API_TOKEN", "123456")
		server := InitServer(t)

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		expectedCode := http.StatusUnauthorized

		assert.Equal(t, expectedCode, res.Code)
		os.Setenv("API_TOKEN", "")
	})

}
func TestProductHandler_Show(t *testing.T) {
	t.Run("should return a product and status code OK", func(t *testing.T) {
		// given
		product1 := domain.Product{
			Id:         1,
			Name:       "product 1",
			Quantity:   new(int),
			Code:       "code1",
			Published:  new(bool),
			Expiration: "2024-10-09",
			Price:      0.0,
		}

		db := map[int]*domain.Product{
			1: &product1,
		}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "1")

		hd.Show(res, req)

		// then

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())

	})
	t.Run("should return a bad request when invalid id", func(t *testing.T) {
		// given

		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "batata")

		hd.Show(res, req)

		// then

		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedErrorMessage := "product id must be a number"
		getResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &getResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, expectedErrorMessage, getResponse.Message)
	})
	t.Run("should return not found when product does not exist", func(t *testing.T) {
		// given
		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "1")

		hd.Show(res, req)

		// then

		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedErrorMessage := fmt.Sprintf("product with id %d not found", 1)
		getResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &getResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, expectedErrorMessage, getResponse.Message)
	})
	t.Run("should return a unauthorized code when a request doesnt have a token", func(t *testing.T) {
		// given
		os.Setenv("API_TOKEN", "123456")
		server := InitServer(t)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		expectedCode := http.StatusUnauthorized

		assert.Equal(t, expectedCode, res.Code)
		os.Setenv("API_TOKEN", "")

	})
}
func TestProductHandler_Create(t *testing.T) {
	t.Run("should return a product and status code Created", func(t *testing.T) {
		// given
		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(`{
				"name": "product 1",
				"quantity": 1,
				"code_value": "code1",
				"is_published": true,
				"expiration": "2024-10-08",
				"price": 20.04
			}`))

		res := httptest.NewRecorder()
		hd.Create(res, req)

		// then

		expectedCode := http.StatusCreated
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		getResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &getResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, false, getResponse.Error)
	})
	t.Run("should return a unauthorized code when a request doesnt have a token", func(t *testing.T) {
		// given
		os.Setenv("API_TOKEN", "123456")
		server := InitServer(t)

		req := httptest.NewRequest(http.MethodPost, "/products", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		expectedCode := http.StatusUnauthorized

		assert.Equal(t, expectedCode, res.Code)
		os.Setenv("API_TOKEN", "")
	})

}
func TestProductHandlerUpdate(t *testing.T) {
	t.Run("should return a product and status code OK", func(t *testing.T) {})
	t.Run("should return a bad request when invalid id", func(t *testing.T) {
		// given

		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%d", 1), strings.NewReader(`{
			"name": "product 1",
			"quantity": 1,
			"code_value": "code1",
			"is_published": true,
			"expiration": "2024-10-08",
			"price": 20.04
		}`))
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "batata")

		hd.Update(res, req)

		// then

		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedErrorMessage := "product id must be a number"
		showResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &showResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, expectedErrorMessage, showResponse.Message)
	})
	t.Run("should return not found when product does not exist", func(t *testing.T) {
		// given
		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "1")

		hd.Update(res, req)

		// then

		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedErrorMessage := fmt.Sprintf("product with id %d not found", 1)
		updateResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &updateResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, expectedErrorMessage, updateResponse.Message)
	})

	t.Run("should return a unauthorized code when a request doesnt have a token", func(t *testing.T) {
		// given
		os.Setenv("API_TOKEN", "123456")
		server := InitServer(t)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		expectedCode := http.StatusUnauthorized

		assert.Equal(t, expectedCode, res.Code)
		os.Setenv("API_TOKEN", "")
	})

}
func TestProductHandlerDelete(t *testing.T) {
	t.Run("should return no content when product is deleted", func(t *testing.T) {
		// given
		product1 := domain.Product{
			Id:         1,
			Name:       "product 1",
			Quantity:   new(int),
			Code:       "code1",
			Published:  new(bool),
			Expiration: "2024-10-09",
			Price:      0.0,
		}

		db := map[int]*domain.Product{
			1: &product1,
		}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "1")

		hd.Delete(res, req)

		// then

		expectedCode := http.StatusNoContent
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		deleteResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &deleteResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, false, deleteResponse.Error)
	})
	t.Run("should return a bad request when invalid id", func(t *testing.T) {
		// given
		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "batata")

		hd.Delete(res, req)

		// then

		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedErrorMessage := "product id must be a number"
		deleteResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &deleteResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, expectedErrorMessage, deleteResponse.Message)
	})
	t.Run("should return not found when product does not exist", func(t *testing.T) {
		// given
		db := map[int]*domain.Product{}
		repository := products.NewRepository(db)
		service := products.NewService(repository)
		hd := handlers.NewProductHandler(service)

		// when (quando)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()
		req = testutils.WithUrlParam(t, req, "productId", "1")

		hd.Delete(res, req)

		// then

		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedErrorMessage := fmt.Sprintf("product with id %d not found", 1)
		deleteResponse := web.Response{}

		err := json.Unmarshal(res.Body.Bytes(), &deleteResponse)
		assert.NoError(t, err)

		assert.Equal(t, expectedCode, res.Code)
		assert.Equal(t, expectedHeader, res.Header())
		assert.Equal(t, expectedErrorMessage, deleteResponse.Message)
	})

	t.Run("should return a unauthorized code when a request doesnt have a token", func(t *testing.T) {
		// given
		os.Setenv("API_TOKEN", "123456")
		server := InitServer(t)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%d", 1), nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		expectedCode := http.StatusUnauthorized

		assert.Equal(t, expectedCode, res.Code)
		os.Setenv("API_TOKEN", "")
	})

}

func InitServer(t *testing.T) *chi.Mux {
	t.Helper()
	r := chi.NewRouter()
	routes := routes.NewRoutes(r)
	routes.MapRoutes(r)
	return r
}
