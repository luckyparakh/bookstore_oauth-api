package access_token

import (
	"bookstore/src/github.com/luckyparakh/bookstore_oauth-api/src/domain/utils/errors"
	"strings"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpireTime(AccessToken) *errors.RestErr
}

// https://globallogic.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go/learn/lecture/16868300#questions/10623680
type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpireTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
	// repository db.DbRepository
}

func NewService(repo Repository) Service {
	// func NewService(repo db.DbRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(at_id string) (*AccessToken, *errors.RestErr) {
	at_id = strings.TrimSpace(at_id)
	if len(at_id) == 0 {
		return nil, errors.NewBadRequestError("Invalid token")
	}
	accessToken, err := s.repository.GetById(at_id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) UpdateExpireTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpireTime(at)
}
func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}
