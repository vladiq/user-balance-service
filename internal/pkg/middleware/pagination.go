package middleware

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
)

type (
	// CustomKey is used to refer to the context key that stores custom values of this api to avoid overwrites
	CustomKey string
)

const (
	// PageIDKey refers to the context key that stores the next page id
	PageIDKey CustomKey = "page_id"
)

// Pagination middleware is used to extract the next page id from the url query
func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PageID := r.URL.Query().Get(string(PageIDKey))

		uuidPageID := uuid.UUID{}
		var err error

		if PageID != "" {
			uuidPageID, err = uuid.Parse(PageID)
			if err != nil {
				render.Status(r, http.StatusBadRequest)
				render.PlainText(w, r, err.Error())
				return
			}
		}

		ctx := context.WithValue(r.Context(), PageIDKey, uuidPageID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
