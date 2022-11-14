package response

import (
	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/domain"
	"net/http"
)

type GetUserTransfers struct {
	Items      []*domain.Transfer `json:"items"`
	NextPageID uuid.UUID          `json:"next_page_id,omitempty"`
}

func (g *GetUserTransfers) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}
