package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igolubevJ/bookstore_users-api/domain/users"
	"github.com/igolubevJ/bookstore_users-api/services"
	"github.com/igolubevJ/bookstore_users-api/utils/errors"
)

// CreateUser - end point create user in DB
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser - end point get user by user id in DB
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
