package main

import (
	"user-jwt-auth/controllers"
	"user-jwt-auth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadConfigVariales()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/user-registration", controllers.SignUp)
	r.POST("/user-login", controllers.Login)

	err := r.Run(":8080")
	if err != nil {
		panic("could not start the server")
	}
}
