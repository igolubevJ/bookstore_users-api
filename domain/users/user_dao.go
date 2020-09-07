package users

import (
	"fmt"
	"time"

	"github.com/igolubevJ/bookstore_users-api/utils/errors"
)

// Mock DB
var (
	usersDB = make(map[int64]*User)
)

// Get - get user by user ID from DB
func (user *User) Get() *errors.RESTErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

// Save - save user to DB
func (user *User) Save() *errors.RESTErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}

	now := time.Now()

	usersDB[user.ID] = user
	return nil
}
