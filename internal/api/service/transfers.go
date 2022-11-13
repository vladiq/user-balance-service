package service

import (
	"context"

	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type transfersRepo interface {
	MakeTransaction(ctx context.Context, entity domain.Transaction) error
	GetUserMonthlyReport(ctx context.Context, entity domain.Transfer) ([]*domain.Transfer, error)
}

type transfers struct {
	repo   transfersRepo
	mapper mapper.Transfer
}

func NewTransfers(repo transfersRepo) *transfers {
	return &transfers{repo: repo, mapper: mapper.Transfer{}}
}

func (s *transfers) MakeTransfer(ctx context.Context, request request.MakeTransfer) error {
	return s.repo.MakeTransaction(ctx, s.mapper.MakeTransferEntity(request))
}

func (s *transfers) UserMonthlyReport(ctx context.Context, request request.UserMonthlyReport) ([]*response.GetUserMonthlyReport, error) {
	entityResults, err := s.repo.GetUserMonthlyReport(ctx, s.mapper.UserMonthlyReport(request))
	if err != nil {
		return nil, err
	}

	var responseResults []*response.GetUserMonthlyReport
	for _, e := range entityResults {
		responseResults = append(responseResults, s.mapper.EntityToReportEntry(*e))
	}

	return responseResults, nil
}
