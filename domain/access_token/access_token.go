package access_token

import (
	"fmt"
	"github.com/nguyenphucthienan/book-store-user-service/util/crypto_util"
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
	"strings"
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

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expire: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (accessToken *AccessToken) Validate() errors.RestError {
	accessToken.AccessToken = strings.TrimSpace(accessToken.AccessToken)
	if accessToken.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token ID")
	}
	if accessToken.UserId <= 0 {
		return errors.NewBadRequestError("Invalid user ID")
	}
	if accessToken.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid client ID")
	}
	if accessToken.Expire <= 0 {
		return errors.NewBadRequestError("Invalid expiration time")
	}
	return nil
}

func (accessToken *AccessToken) IsExpired() bool {
	return time.Unix(accessToken.Expire, 0).Before(time.Now().UTC())
}

func (accessToken *AccessToken) Generate() {
	accessToken.AccessToken = crypto_util.GetMd5(fmt.Sprintf("at-%d-%d-ran", accessToken.UserId, accessToken.Expire))
}
