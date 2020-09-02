package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartAppication - func starup bookstore_users-API
func StartAppication() {
	mapUrls()
	router.Run(":8080")
}
