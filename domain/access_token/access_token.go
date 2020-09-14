package access_token

import (
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expire      int64  `json:"expire"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expire: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expire, 0).Before(time.Now().UTC())
}
