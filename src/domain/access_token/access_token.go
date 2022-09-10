package access_token

import "time"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    string `json:"client_id"`
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
