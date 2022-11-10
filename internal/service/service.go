package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Repository interface {
	AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error
	MakeReservation(ctx context.Context, userID uuid.UUID, orderID uuid.UUID, amount float64) error
	AcceptReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
}

type Service interface {
	AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error
	MakeReservation(ctx context.Context, userID uuid.UUID, orderID uuid.UUID, amount float64) error
	AcceptReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
}

type BalanceService struct {
	logger zerolog.Logger
	repo   Repository
}

func NewBalanceService(logger zerolog.Logger, repo Repository) *BalanceService {
	return &BalanceService{
		logger: logger,
		repo:   repo,
	}
}

func (bs *BalanceService) AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (bs *BalanceService) MakeReservation(ctx context.Context, userID uuid.UUID, orderID uuid.UUID, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (bs *BalanceService) AcceptReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (bs *BalanceService) GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	return bs.repo.GetUserBalance(ctx, userID)
}
