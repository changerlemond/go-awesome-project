package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}
