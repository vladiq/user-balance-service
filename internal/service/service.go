package service

import (
	"context"

	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type balanceService struct {
	logger zerolog.Logger

	accountRepo     accountRepo
	reservationRepo reservationRepo
	transactionRepo transactionRepo
}

func NewBalanceService(logger zerolog.Logger, ar accountRepo, rr reservationRepo, tr transactionRepo) *balanceService {
	return &balanceService{
		logger:          logger,
		accountRepo:     ar,
		reservationRepo: rr,
		transactionRepo: tr,
	}
}

func (bs *balanceService) GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	return bs.accountRepo.GetUserBalance(ctx, userID)
}
