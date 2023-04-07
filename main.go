package main

import (
	"user-jwt-auth/controllers"
	"user-jwt-auth/initializers"
	"user-jwt-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadConfigVariales()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/user-registration", controllers.SignUpHandler)
	r.POST("/user-login", controllers.LoginHandler)
	r.GET("/user-details", middlewares.ReqAuthorization, controllers.UserDetailsHandler)

	err := r.Run(":8080")
	if err != nil {
		panic("could not start the server")
	}
}
