package services

import (
	"github.com/igolubevJ/bookstore_users-api/domain/users"
	"github.com/igolubevJ/bookstore_users-api/utils/errors"
)

// CreateUser - func create new user
func CreateUser(user users.User) (*users.User, *errors.RESTErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return &user, nil
}
