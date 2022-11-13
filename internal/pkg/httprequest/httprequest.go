package httprequest

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

func NewRequest(
	method string,
	target string,
	body string,
	query map[string]string,
	path map[string]string,
) *http.Request {
	req := httptest.NewRequest(method, target, bytes.NewBufferString(body))

	rctx := chi.NewRouteContext()
	for k, v := range path {
		rctx.URLParams.Add(k, v)
	}

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	return req
}
