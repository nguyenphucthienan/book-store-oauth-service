package service

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/nguyenphucthienan/book-store-oauth-service/domain/user"
	"github.com/nguyenphucthienan/book-store-oauth-service/utils/errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082/api",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserService interface {
	LoginUser(string, string) (*user.User, *errors.RestError)
}

type restUserService struct{}

func NewRestUserService() RestUserService {
	return &restUserService{}
}

func (s *restUserService) LoginUser(email string, password string) (*user.User, *errors.RestError) {
	request := user.LoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		restErr, err := errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to login user")
		}
		return nil, restErr
	}

	var returnedUser user.User
	if err := json.Unmarshal(response.Bytes(), &returnedUser); err != nil {
		return nil, errors.NewInternalServerError("Error when trying to unmarshal user login response")
	}

	return &returnedUser, nil
}
