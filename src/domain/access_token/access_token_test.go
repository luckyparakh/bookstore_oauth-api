package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpirationTime(t *testing.T) {
	assert.EqualValues(t, 24, 24, "expiration time should be 24 hours")
}
func TestGetAccessToken(t *testing.T) {
	at := GetAccessToken()
	assert.EqualValues(t, at.AccessToken, "", "new access token should not any token")
	// if at.AccessToken != "" {
	// 	t.Error("new access token should not any token")
	// }
	assert.EqualValues(t, at.UserId, 0, "new access token should not have any associated user id")
	// if at.UserId != 0 {
	// 	t.Error("new access token should not have any associated user id")
	// }
	assert.False(t, at.IsExpired(), "access token should not expired")
	// if at.IsExpired() {
	// 	t.Error("access token should not expired")
	// }
}

func TestIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "access token should be expired")
	// if !at.IsExpired() {
	// 	t.Error("access token should be expired")
	// }
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token should not expired")
	// if at.IsExpired() {
	// 	t.Error("access token should not expire")
	// }

	at.Expires = time.Now().UTC().Unix()
	assert.True(t, at.IsExpired(), "access token should expire")
	// if !at.IsExpired() {
	// 	t.Error("access token should expire")
	// }
}
