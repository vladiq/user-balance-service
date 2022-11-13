package service

import (
	"context"
	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type transfersRepo interface {
	MakeTransaction(ctx context.Context, entity domain.Transaction) error
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
