package service

import (
	"context"

	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type reservationsRepo interface {
	Create(ctx context.Context, entity domain.Reservation) error
	Delete(ctx context.Context, entity domain.Reservation) error
}

type reservations struct {
	repo   reservationsRepo
	mapper mapper.Reservation
}

func NewReservations(repo reservationsRepo) *reservations {
	return &reservations{repo: repo, mapper: mapper.Reservation{}}
}

func (s *reservations) CreateReservation(ctx context.Context, request request.CreateReservation) error {
	return s.repo.Create(ctx, s.mapper.MakeCreateReservationEntity(request))
}

func (s *reservations) CancelReservation(ctx context.Context, request request.CancelReservation) error {
	return s.repo.Delete(ctx, s.mapper.MakeCancelReservationEntity(request))
}
