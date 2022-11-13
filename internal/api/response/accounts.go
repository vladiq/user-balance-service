package response

import "github.com/google/uuid"

// Возвращаемые жейсоны. Заполняются в маппере.
type GetAccount struct {
	ID      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}
