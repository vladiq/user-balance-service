package service

import (
	"context"
	"fmt"

	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type balanceService struct {
	logger zerolog.Logger

	accountRepo     AccountRepo
	reservationRepo ReservationRepo
	transferRepo    TransferRepo
}

func NewBalanceService(logger zerolog.Logger, ar AccountRepo, rr ReservationRepo, tr TransferRepo) *balanceService {
	return &balanceService{
		logger:          logger,
		accountRepo:     ar,
		reservationRepo: rr,
		transferRepo:    tr,
	}
}

func (bs *balanceService) GetAccount(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	account, err := bs.accountRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("getting user balance: %w", err)
	}
	return account, nil
}

func (bs *balanceService) CreateAccount(ctx context.Context, amount float64) error {
	err := bs.accountRepo.CreateAccount(ctx, amount)
	if err != nil {
		return fmt.Errorf("creating account: %w", err)
	}

	return nil
}

func (bs *balanceService) UpdateBalance(ctx context.Context, userID uuid.UUID, amount float64) error {
	switch amount >= 0 {
	case true:
		err := bs.accountRepo.DepositFunds(ctx, userID, amount)
		if err != nil {
			return fmt.Errorf("depositing funds to user account: %w", err)
		}
	default:
		err := bs.accountRepo.WithdrawFunds(ctx, userID, -amount)
		if err != nil {
			return fmt.Errorf("withdrawing funds from user account: %w", err)
		}
	}

	return nil
}

//
//func (bs *balanceService) CreateReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error {
//
//}
