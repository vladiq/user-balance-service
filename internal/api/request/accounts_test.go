package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vladiq/user-balance-service/internal/pkg/httprequest"
	"github.com/vladiq/user-balance-service/internal/testdata"
)

func TestCreateAccountBind(t *testing.T) {
	expected := CreateAccount{
		Amount: 5.,
	}

	var actual CreateAccount

	err := actual.Bind(httprequest.NewRequest(
		http.MethodPost,
		"/account",
		testdata.CreateAccountValid,
		nil,
		nil))

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
			body:     testdata.CreateAccountInvalidJson,
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
			require.EqualError(t, r.Bind(httprequest.NewRequest(http.MethodPost,
				"/account",
				tc.body,
				nil,
				nil)), tc.errorMsg)
		})
	}
}
