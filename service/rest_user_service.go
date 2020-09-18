package service

import (
	"encoding/json"
	"errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/nguyenphucthienan/book-store-oauth-service/domain/user"
	restErrors "github.com/nguyenphucthienan/book-store-utils-go/errors"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082/api",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserService interface {
	LoginUser(string, string) (*user.User, restErrors.RestError)
}

type restUserService struct{}

func NewRestUserService() RestUserService {
	return &restUserService{}
}

func (s *restUserService) LoginUser(email string, password string) (*user.User, restErrors.RestError) {
	request := user.LoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, restErrors.NewInternalServerError("Invalid rest client response when trying to login user",
			errors.New("rest client error"))
	}

	if response.StatusCode > 299 {
		restErr, err := restErrors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, restErrors.NewInternalServerError("Invalid error interface when trying to login user", restErr)
		}
		return nil, restErr
	}

	var returnedUser user.User
	if err := json.Unmarshal(response.Bytes(), &returnedUser); err != nil {
		return nil, restErrors.NewInternalServerError("Error when trying to unmarshal user login response", err)
	}

	return &returnedUser, nil
}
