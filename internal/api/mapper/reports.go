package mapper

import (
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Report struct {
}

func (r Report) GetResponseReportEntries(reports []*domain.Report) []*response.GetServiceReport {
	responseEntries := make([]*response.GetServiceReport, len(reports))

	for i, report := range reports {
		responseEntry := response.GetServiceReport{
			ServiceID: report.ServiceID,
			Amount:    report.Amount,
		}
		responseEntries[i] = &responseEntry
	}

	return responseEntries
}
