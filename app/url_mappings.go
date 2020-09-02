package app

import "github.com/igolubevJ/bookstore_users-api/controllers/users"

func mapUrls() {
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
