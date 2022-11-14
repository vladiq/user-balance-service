package request

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"strconv"
)

type GetServiceReport struct {
	Month int
}

func (r *GetServiceReport) Bind(req *http.Request) error {
	month, err := strconv.Atoi(req.URL.Query().Get("month"))
	if err != nil {
		return fmt.Errorf("converting month to number: %w", err)
	}

	r.Month = month

	return r.validate()
}

func (r *GetServiceReport) validate() error {
	fmt.Println(r.Month)

	if err := validation.Validate(
		r.Month,
		validation.Required,
		validation.Min(1),
		validation.Max(12),
	); err != nil {
		return fmt.Errorf("validating month: %w", err)
	}

	return nil
}
