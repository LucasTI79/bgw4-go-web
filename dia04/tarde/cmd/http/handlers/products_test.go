package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lucasti79/bgw4-put-patch-delete/cmd/http/handlers"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/storage"
	"github.com/lucasti79/bgw4-put-patch-delete/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductsHandler_Get(t *testing.T) {

	t.Run("should return status 200 and a list of products", func(t *testing.T) {
		// given (dado)
		db := map[int]storage.ProductAttributes{
			1: {Name: "product 1", Type: "type 1", Quantity: 1, Price: 1.1},
			2: {Name: "product 2", Type: "type 2", Quantity: 2, Price: 2.2},
		}
		st := storage.NewProductsStorage(db)
		hd := handlers.NewProductsHandler(st)

		// when (quando)

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()
		hd.Get()(res, req)

		// then (entao)

		expectedCode := http.StatusOK
		expectedData := []domain.Product{
			{
				Id:       1,
				Name:     "product 1",
				Type:     "type 1",
				Quantity: 1,
				Price:    1.1,
			}, {
				Id:       2,
				Name:     "product 2",
				Type:     "type 2",
				Quantity: 2,
				Price:    2.2,
			},
		}

		var getResponse struct {
			Message string            `json:"message"`
			Data    *[]domain.Product `json:"data"`
			Error   bool              `json:"error"`
		}

		err := json.Unmarshal(res.Body.Bytes(), &getResponse)
		assert.NoError(t, err)

		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedHeader, res.Header())

		require.ElementsMatch(t, expectedData, *getResponse.Data)

	})
}

func TestProductsHandler_Save(t *testing.T) {
	t.Run("should return status 200 and a product", func(t *testing.T) {
		// given
		db := map[int]storage.ProductAttributes{
			1: {Name: "product 1", Type: "type 1", Quantity: 1, Price: 1.1},
		}
		st := storage.NewProductsStorage(db)
		hd := handlers.NewProductsHandler(st)

		// when
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%s", "1"), strings.NewReader(
			`{"name": "product 1", "type": "type 1", "count": 1, "price": 1.1}`,
		))

		req = testutils.WithUrlParam(t, req, "id", "1")

		res := httptest.NewRecorder()
		hd.UpdateOrCreate()(res, req)

		// then
		expectedCode := http.StatusOK
		expectedData := domain.Product{
			Id: 1, Name: "product 1", Type: "type 1", Quantity: 1, Price: 1.1,
		}
		var saveResponse struct {
			Message string          `json:"message"`
			Data    *domain.Product `json:"data"`
			Error   bool            `json:"error"`
		}

		err := json.Unmarshal(res.Body.Bytes(), &saveResponse)
		assert.NoError(t, err)
		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, res.Header())
		assert.ObjectsAreEqualValues(expectedData, *saveResponse.Data)
	})
}
