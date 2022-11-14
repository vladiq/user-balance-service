package request

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vladiq/user-balance-service/internal/pkg/httprequest"
	"github.com/vladiq/user-balance-service/internal/testdata"
	"net/http"
	"testing"
)

func TestCreateReservationBind(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := CreateReservation{
		validUUID,
		validUUID,
		validUUID,
		5.5,
	}

	var actual CreateReservation
	err := actual.Bind(
		httprequest.NewRequest(
			http.MethodPost,
			"/reservations",
			testdata.CreateReservationValid,
			nil,
			nil,
		),
	)

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestCreateReservationBind_error(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		errorMsg string
	}{
		{
			name:     "invalid json",
			body:     testdata.CreateReservationInvalidJSON,
			errorMsg: "binding body: json: cannot unmarshal string into Go value of type request.CreateReservation",
		},
		{
			name:     "negative amount",
			body:     testdata.CreateReservationNegativeAmount,
			errorMsg: "validating body: amount: must be no less than 0.",
		},
		{
			name:     "invalid UUID",
			body:     testdata.CreateReservationInvalidUUID,
			errorMsg: "binding body: invalid UUID length: 3",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var r CreateReservation
			require.EqualError(
				t,
				r.Bind(httprequest.NewRequest(http.MethodPost, "/reservations", tc.body, nil, nil)),
				tc.errorMsg,
			)
		})
	}
}
