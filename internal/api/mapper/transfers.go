package mapper

import (
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Transfer struct {
}

func (m Transfer) MakeTransferEntity(r request.MakeTransfer) domain.Transaction {
	return domain.Transaction{
		FromID: r.FromID,
		ToID:   r.ToID,
		Amount: r.Amount,
	}
}
