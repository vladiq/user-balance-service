package mapper

import (
	"github.com/google/uuid"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
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

func (m Transfer) UserMonthlyReport(r request.GetUserTransfers) domain.Transfer {
	return domain.Transfer{
		AccountID: r.AccountID,
	}
}

//func (m Transfer) EntityToReportEntry(e domain.Transfer) *response.GetUserMonthlyReport {
//	return &response.GetUserMonthlyReport{
//		Timestamp: e.CreatedAt,
//		IsAccrual: e.IsAccrual,
//		Info:      e.Info,
//		Amount:    e.Amount,
//	}
//}

func (m Transfer) GetResponseMonthlyReport(transfers []*domain.Transfer, nextPageID uuid.UUID) response.GetUserTransfers {
	return response.GetUserTransfers{
		Items:      transfers,
		NextPageID: nextPageID,
	}
}
