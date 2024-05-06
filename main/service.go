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

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")

	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var updateUser User
	if err := c.BindJSON(&updateUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updateFields := make(map[string]interface{})
	if updateUser.Name != "" {
		updateFields["name"] = updateUser.Name
	}
	if updateUser.Email != "" {
		updateFields["email"] = updateUser.Email
	}
	if updateUser.Password != "" {
		updateFields["password"] = updateUser.Password
	}

	if err := db.Model(&user).Updates(updateFields).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var updatedUser User
	if err := db.First(&updatedUser, id).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		db.Delete(&user)
	}
}
