package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() // default setting
	// 1. String data
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World") // http status, response data
	}) // url, handler function
	// 2. JSON data
	r.GET("/json", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "Hello World",
		}) // gin.H{}: map you can configure
	}) // url, handler function
	r.Run("localhost:8080") // api running url, port
}
