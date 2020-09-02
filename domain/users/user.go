package users

import (
	"strings"

	"github.com/igolubevJ/bookstore_users-api/utils/errors"
)

// User - user entity struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate - User empty email validator
func (user *User) Validate() *errors.RESTErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
