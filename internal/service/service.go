package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type Repository interface {
	AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error
	GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error)
}

type Service struct {
	repo   Repository
	logger zerolog.Logger
}

func NewService(repo Repository, logger zerolog.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) GetUserBalance(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	account, err := s.repo.GetUserBalance(ctx, userID)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *Service) AddFundsToAccount(ctx context.Context, userID uuid.UUID, amount float64) error {
	err := s.repo.AddFundsToAccount(ctx, userID, amount)
	return err
}
