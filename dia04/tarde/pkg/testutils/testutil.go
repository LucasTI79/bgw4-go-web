package testutils

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

func WithUrlParam(t *testing.T, r *http.Request, key, value string) *http.Request {
	t.Helper()
	chiCtx := chi.NewRouteContext()
	req := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
	chiCtx.URLParams.Add(key, value)

	return req
}

// WithUrlParams returns a pointer to a request object with the given URL params
// added to a new chi.Context object. for single param assignment see WithUrlParam
func WithUrlParamst(t *testing.T, r *http.Request, params map[string]string) *http.Request {
	t.Helper()

	chiCtx := chi.NewRouteContext()
	req := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))
	for key, value := range params {
		chiCtx.URLParams.Add(key, value)
	}
	return req
}
