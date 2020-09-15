package repository

import (
	"github.com/gocql/gocql"
	"github.com/nguyenphucthienan/book-store-oauth-service/clients/cassandra"
	"github.com/nguyenphucthienan/book-store-oauth-service/domain/access_token"
	"github.com/nguyenphucthienan/book-store-oauth-service/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expire FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expire) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expire = ? WHERE access_token = ?;"
)

func NewAccessTokenRepository() AccessTokenRepository {
	return &accessTokenRepository{}
}

type AccessTokenRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type accessTokenRepository struct{}

func (r *accessTokenRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	var accessToken access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&accessToken.AccessToken,
		&accessToken.UserId,
		&accessToken.ClientId,
		&accessToken.Expire,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token found with given ID")
		}
		return nil, errors.NewInternalServerError("Error when trying to get current ID")
	}
	return &accessToken, nil
}

func (r *accessTokenRepository) Create(accessToken access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		accessToken.AccessToken,
		accessToken.UserId,
		accessToken.ClientId,
		accessToken.Expire,
	).Exec(); err != nil {
		return errors.NewInternalServerError("Error when trying to save access token in database")
	}
	return nil
}

func (r *accessTokenRepository) UpdateExpirationTime(accessToken access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		accessToken.Expire,
		accessToken.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("Error when trying to update current resource")
	}
	return nil
}
