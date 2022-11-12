package service

import (
	"context"

	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
)

type AccountRepo interface {
	GetUser(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
	CreateAccount(ctx context.Context, amount float64) error
	DepositFunds(ctx context.Context, userID uuid.UUID, amount float64) error
	WithdrawFunds(ctx context.Context, userID uuid.UUID, amount float64) error
}

type ReservationRepo interface {
	CreateReservation(ctx context.Context, userID uuid.UUID, serviceID uuid.UUID, orderID uuid.UUID, amount float64) error
}

type TransferRepo interface {
}
