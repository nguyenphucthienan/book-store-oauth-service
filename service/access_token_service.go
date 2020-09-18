package service

import (
	"github.com/nguyenphucthienan/book-store-oauth-service/domain/access_token"
	"github.com/nguyenphucthienan/book-store-oauth-service/repository"
	"github.com/nguyenphucthienan/book-store-utils-go/errors"
	"strings"
)

func NewAccessTokenService(
	accessTokenRepository repository.AccessTokenRepository,
	restUserService RestUserService,
) AccessTokenService {
	return &accessTokenService{
		accessTokenRepository: accessTokenRepository,
		restUserService:       restUserService,
	}
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, errors.RestError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, errors.RestError)
	UpdateExpirationTime(access_token.AccessToken) errors.RestError
}

type accessTokenService struct {
	accessTokenRepository repository.AccessTokenRepository
	restUserService       RestUserService
}

func (s *accessTokenService) GetById(accessTokenId string) (*access_token.AccessToken, errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token ID")
	}
	accessToken, err := s.accessTokenRepository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *accessTokenService) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken,
	errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the User Service's API
	user, err := s.restUserService.LoginUser(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token
	accessToken := access_token.GetNewAccessToken(user.Id)
	accessToken.Generate()

	// Save the new access token in Cassandra
	if err := s.accessTokenRepository.Create(accessToken); err != nil {
		return nil, err
	}

	return &accessToken, nil
}

func (s *accessTokenService) UpdateExpirationTime(accessToken access_token.AccessToken) errors.RestError {
	if err := accessToken.Validate(); err != nil {
		return err
	}
	return s.accessTokenRepository.UpdateExpirationTime(accessToken)
}
