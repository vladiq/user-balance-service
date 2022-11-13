package constant

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
)
