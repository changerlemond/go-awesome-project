package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	connectToDatabase()
	r := gin.Default() // default setting
	r.GET("/ping", func(c *gin.Context) {})
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)
	r.Run("localhost:8080") // api running url, port
}
