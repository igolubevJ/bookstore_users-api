package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser - end point create user in DB
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

// GetUser - end point get user by user id in DB
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
