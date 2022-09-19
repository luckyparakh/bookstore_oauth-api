package db

import (
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/clients/cassandra"
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/domain/access_token"
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/domain/utils/errors"

	"github.com/gocql/gocql"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpireTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

const (
	queryGetAccessToken    = "Select access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	queryInsertAccessToken = "Insert into access_tokens (access_token, user_id, client_id, expires) values (?,?,?,?);"
	queryUpdateExpireTime  = "Update access_tokens Set expires=? where access_token=?;"
)

func (d *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken
	if err := cassandra.GetNewSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken, &result.UserId,
		&result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.InternalServerError("token not found")
		}
		return nil, errors.InternalServerError(err.Error())
	}

	return &result, nil
}
func (d *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetNewSession().Query(queryInsertAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

func (d *dbRepository) UpdateExpireTime(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetNewSession().Query(queryUpdateExpireTime, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
