package access_token

import (
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/domain/utils/errors"
	"strings"
	"time"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

const expirationTime = 24

func GetAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}
func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
	// Another way
	// timeNow := time.Now().UTC().Unix()
	// return timeNow >= at.Expires
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid exipre time")
	}
	return nil
}
