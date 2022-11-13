package mapper

import (
	"time"

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

func (m Transfer) UserMonthlyReport(r request.UserMonthlyReport) domain.Transfer {
	return domain.Transfer{
		AccountID: r.AccountID,
		CreatedAt: time.Date(r.Year, time.Month(r.Month), 1, 0, 0, 0, 0, time.UTC),
	}
}

func (m Transfer) EntityToReportEntry(e domain.Transfer) *response.GetUserMonthlyReport {
	return &response.GetUserMonthlyReport{
		Timestamp: e.CreatedAt,
		IsAccrual: e.IsAccrual,
		Info:      e.Info,
		Amount:    e.Amount,
	}
}
