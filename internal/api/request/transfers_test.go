package request

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/vladiq/user-balance-service/internal/pkg/httprequest"
	"github.com/vladiq/user-balance-service/internal/testdata"
	"net/http"
	"testing"
)

func TestMakeTransferBind(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)
	expected := MakeTransfer{
		validUUID,
		validUUID,
		5.5,
	}

	var actual MakeTransfer
	err := actual.Bind(
		httprequest.NewRequest(
			http.MethodPost,
			"/transfers",
			testdata.MakeTransferValid,
			nil,
			nil,
		),
	)

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestMakeTransferBind_error(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		errorMsg string
	}{
		{
			name:     "invalid json",
			body:     testdata.MakeTransferInvalidJSON,
			errorMsg: "binding body: unexpected EOF",
		},
		{
			name:     "negative amount",
			body:     testdata.MakeTransferNegativeAmount,
			errorMsg: "validating amount: must be no less than 0",
		},
		{
			name:     "invalid UUID",
			body:     testdata.MakeTransferInvalidUUID,
			errorMsg: "binding body: invalid UUID length: 3",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var r MakeTransfer
			require.EqualError(
				t,
				r.Bind(httprequest.NewRequest(http.MethodPost, "/transfers", tc.body, nil, nil)),
				tc.errorMsg,
			)
		})
	}
}
