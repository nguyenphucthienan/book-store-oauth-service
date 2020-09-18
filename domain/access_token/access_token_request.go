package access_token

import (
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
)

const (
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Email    string `json:"email"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (accessTokenRequest *AccessTokenRequest) Validate() errors.RestError {
	switch accessTokenRequest.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("Invalid grant_type parameter")
	}

	// TODO: Validate parameters for each grant type
	return nil
}
