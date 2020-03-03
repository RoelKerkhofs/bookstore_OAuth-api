package db

import (
	"bookstore/bookstore_OAuth-api/src/clients/cassandra"
	"bookstore/bookstore_OAuth-api/src/domain/access_token"
	"bookstore/bookstore_OAuth-api/src/utils/errors"
	"fmt"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken       = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken    = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES ( ?, ?, ?, ?);"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.Expires,
		&result.ClientId,
	); err != nil {

		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	fmt.Println("Ophehaalde token: ", result.AccessToken)
	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.Expires,
		token.ClientId,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpirationTime, token.Expires, token.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
