package services

import (
	"github.com/igolubevJ/bookstore_users-api/domain/users"
	"github.com/igolubevJ/bookstore_users-api/utils/errors"
)

// CreateUser - create new user
func CreateUser(user users.User) (*users.User, *errors.RESTErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser - get user by id
func GetUser(userID int64) (*users.User, *errors.RESTErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
