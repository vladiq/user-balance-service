package mapper

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vladiq/user-balance-service/internal/api/response"
	"github.com/vladiq/user-balance-service/internal/domain"
	"github.com/vladiq/user-balance-service/internal/testdata"
	"testing"
)

func TestGetResponseReportEntries(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)

	expected := []*response.GetServiceReport{
		{
			ServiceID: validUUID,
			Amount:    5.5,
		},
	}

	entities := []*domain.Report{
		{
			ServiceID: validUUID,
			Amount:    5.5,
		},
	}

	require.Equal(t, expected, Report{}.GetResponseReportEntries(entities))
}
