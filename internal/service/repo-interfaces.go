package service

import (
	"context"

	"github.com/vladiq/user-balance-service/internal/domain"

	"github.com/google/uuid"
)

type accountRepo interface {
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
}

type reservationRepo interface {
}

type transactionRepo interface {
}
