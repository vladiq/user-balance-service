package response

import "github.com/google/uuid"

type GetServiceReport struct {
	ServiceID uuid.UUID
	Amount    float64
}
