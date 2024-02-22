package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vaibhavsingh9/go-jwt/controllers"
	"github.com/vaibhavsingh9/go-jwt/initializers"
	"github.com/vaibhavsingh9/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
