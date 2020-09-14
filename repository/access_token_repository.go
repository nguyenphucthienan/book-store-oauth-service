package repository

import (
	"github.com/nguyenphucthienan/book-store-oauth-service/domain/access_token"
	"github.com/nguyenphucthienan/book-store-oauth-service/utils/errors"
)

func NewRepository() AccessTokenRepository {
	return &accessTokenRepository{}
}

type AccessTokenRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type accessTokenRepository struct{}

func (r *accessTokenRepository) GetById(string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("Database connection not implemented")
}
