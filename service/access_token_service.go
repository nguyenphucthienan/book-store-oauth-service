package service

import (
	"github.com/nguyenphucthienan/book-store-oauth-service/domain/access_token"
	"github.com/nguyenphucthienan/book-store-oauth-service/repository"
	"github.com/nguyenphucthienan/book-store-oauth-service/utils/errors"
	"strings"
)

func NewService(repository repository.AccessTokenRepository) AccessTokenService {
	return &accessTokenService{
		repository: repository,
	}
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type accessTokenService struct {
	repository repository.AccessTokenRepository
}

func (s *accessTokenService) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Invalid access token ID")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
