package utils

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/vladiq/user-balance-service/internal/constant"
)

func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, constant.ErrBadRequest) {
		render.Status(r, http.StatusBadRequest)
	} else if errors.Is(err, constant.ErrNotFound) {
		render.Status(r, http.StatusNotFound)
	} else {
		render.Status(r, http.StatusInternalServerError)
	}
	render.PlainText(w, r, err.Error())
}
