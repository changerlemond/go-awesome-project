package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	connectToDatabase()
	router := gin.Default() // default setting
	protected := router.Group("/")
	protected.Use(MiddlewareJWTAuth())
	{
		protected.GET("/users", GetUsers)
		protected.GET("/users/:id", GetUser)
		protected.PUT("/users/:id", UpdateUser)
		protected.DELETE("/users/:id", DeleteUser)
	}
	router.POST("/sign-up", SignUp)
	router.POST("/login", Login)
	router.Run("localhost:8080") // api running url, port
}
