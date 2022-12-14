package service

import (
	"context"

	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type reservationsRepo interface {
	Create(ctx context.Context, entity domain.Reservation) error
	Cancel(ctx context.Context, entity domain.Reservation) error
	Confirm(ctx context.Context, entity domain.Reservation) error
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
	return s.repo.Cancel(ctx, s.mapper.MakeCancelReservationEntity(request))
}

func (s *reservations) ConfirmReservation(ctx context.Context, request request.ConfirmReservation) error {
	return s.repo.Confirm(ctx, s.mapper.MakeConfirmReservationEntity(request))
}
