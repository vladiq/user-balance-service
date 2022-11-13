package response

import "github.com/google/uuid"

type GetAccount struct {
	ID      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}
