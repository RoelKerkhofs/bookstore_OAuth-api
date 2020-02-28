package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationtime, "Brand new expiration time should be 24")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.NotNil(t, true, at.IsExpired(), "Brand new access token should not be nil")
	assert.EqualValues(t, "", at.AccessToken, "New access token should be empty")
	assert.Empty(t, at.UserId, "New access token should not have an associated id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token should be valid until 3 hours from now")
}
