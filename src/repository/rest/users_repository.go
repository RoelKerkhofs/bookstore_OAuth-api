package rest

import (
	"bookstore/bookstore_OAuth-api/src/domain/users"
	"bookstore/bookstore_OAuth-api/src/utils/errors"
	_ "github.com/golang/protobuf/proto"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	return nil, nil
}
