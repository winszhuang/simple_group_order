package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJWT(t *testing.T) {
	data := map[string]string{
		"userId":    "1",
		"channelId": "15",
	}
	token, err := CreateJWT(data)
	require.NoError(t, err)

	claims, err := VerifyJWT(token)
	require.NoError(t, err)

	require.Equal(t, data["userId"], claims["userId"])
	require.Equal(t, data["channelId"], claims["channelId"])
}
