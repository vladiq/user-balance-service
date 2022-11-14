package service

import (
	"context"
	"github.com/vladiq/user-balance-service/internal/api/mapper"
	"github.com/vladiq/user-balance-service/internal/api/request"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
)

type reportsRepo interface {
	GetReport(ctx context.Context, month int) ([]*domain.Report, error)
}

type reports struct {
	repo   reportsRepo
	mapper mapper.Report
}

func NewReports(repo reportsRepo) *reports {
	return &reports{repo: repo, mapper: mapper.Report{}}
}

func (r *reports) GetServiceReport(ctx context.Context, request request.GetServiceReport) ([]*response.GetServiceReport, error) {
	entries, err := r.repo.GetReport(ctx, request.Month)
	if err != nil {
		return nil, err
	}

	responseEntries := r.mapper.GetResponseReportEntries(entries)
	return responseEntries, nil
}
