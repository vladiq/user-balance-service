package request

import (
	"github.com/google/uuid"
	"net/http"
	"testing"

	"github.com/vladiq/user-balance-service/internal/pkg/httprequest"
	"github.com/vladiq/user-balance-service/internal/testdata"

	"github.com/stretchr/testify/require"
)

func TestCreateAccountBind(t *testing.T) {
	expected := CreateAccount{
		Amount: 5.,
	}

	var actual CreateAccount
	err := actual.Bind(
		httprequest.NewRequest(
			http.MethodPost,
			"/account",
			testdata.CreateAccountValid,
			nil,
			nil,
		),
	)

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestCreateAccountBind_error(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		errorMsg string
	}{
		{
			name:     "invalid json",
			body:     testdata.CreateAccountInvalidJSON,
			errorMsg: "binding body: unexpected EOF",
		},
		{
			name:     "negative amount",
			body:     testdata.CreateAccountNegativeAmount,
			errorMsg: "validating amount: must be no less than 0",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var r CreateAccount
			require.EqualError(
				t,
				r.Bind(httprequest.NewRequest(http.MethodPost, "/account", tc.body, nil, nil)),
				tc.errorMsg,
			)
		})
	}
}

func TestDepositFundsBind(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)

	expected := DepositFunds{
		ID:     validUUID,
		Amount: 5.,
	}

	var actual DepositFunds
	err := actual.Bind(
		httprequest.NewRequest(
			http.MethodPost,
			"/account",
			testdata.DepositFundsValid,
			nil,
			nil,
		),
	)

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestDepositFundsBind_error(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		errorMsg string
	}{
		{
			name:     "invalid json",
			body:     testdata.DepositFundsInvalidJSON,
			errorMsg: "binding body: invalid character 'i' looking for beginning of object key string",
		},
		{
			name:     "negative amount",
			body:     testdata.DepositFundsNegativeAmount,
			errorMsg: "validating amount: must be no less than 0",
		},
		{
			name:     "invalid UUID",
			body:     testdata.DepositFundsInvalidUUID,
			errorMsg: "binding body: json: cannot unmarshal string into Go value of type request.DepositFunds",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var r DepositFunds
			require.EqualError(
				t,
				r.Bind(httprequest.NewRequest(http.MethodPost, "/account", tc.body, nil, nil)),
				tc.errorMsg,
			)
		})
	}
}

func TestWithdrawFundsBind(t *testing.T) {
	validUUID, _ := uuid.Parse(testdata.ValidUUIDString)

	expected := WithdrawFunds{
		ID:     validUUID,
		Amount: 5.,
	}

	var actual WithdrawFunds
	err := actual.Bind(
		httprequest.NewRequest(
			http.MethodPost,
			"/account",
			testdata.WithdrawFundsValid,
			nil,
			nil,
		),
	)

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestWithdrawFundsBind_error(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		errorMsg string
	}{
		{
			name:     "invalid json",
			body:     testdata.WithdrawFundsInvalidJSON,
			errorMsg: "binding body: invalid character 'i' looking for beginning of object key string",
		},
		{
			name:     "negative amount",
			body:     testdata.WithdrawFundsNegativeAmount,
			errorMsg: "validating amount: must be no less than 0",
		},
		{
			name:     "invalid UUID",
			body:     testdata.WithdrawFundsInvalidUUID,
			errorMsg: "binding body: json: cannot unmarshal string into Go value of type request.WithdrawFunds",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var r WithdrawFunds
			require.EqualError(
				t,
				r.Bind(httprequest.NewRequest(http.MethodPost, "/account", tc.body, nil, nil)),
				tc.errorMsg,
			)
		})
	}
}
