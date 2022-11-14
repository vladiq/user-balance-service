package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type transfersRepo interface {
	CreateTransfer(ctx context.Context, entity domain.Transaction) error
	GetTransfers(ctx context.Context, entity domain.Transfer, pageID uuid.UUID, orderBy string) ([]*domain.Transfer, uuid.UUID, error)
}

type transfers struct {
	repo   transfersRepo
	mapper mapper.Transfer
}

func NewTransfers(repo transfersRepo) *transfers {
	return &transfers{repo: repo, mapper: mapper.Transfer{}}
}

func (s *transfers) MakeTransfer(ctx context.Context, r request.MakeTransfer) error {
	return s.repo.CreateTransfer(ctx, s.mapper.MakeTransferEntity(r))
}

func (s *transfers) GetTransfers(ctx context.Context, r request.GetUserTransfers, pageID uuid.UUID) (response.GetUserTransfers, error) {
	orderBy := r.OrderBy
	entityResults, nextPageID, err := s.repo.GetTransfers(ctx, s.mapper.UserMonthlyReport(r), pageID, orderBy)
	if err != nil {
		return response.GetUserTransfers{}, err
	}

	responseResults := s.mapper.GetResponseMonthlyReport(entityResults, nextPageID)
	return responseResults, nil
}
