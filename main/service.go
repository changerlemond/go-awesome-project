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

func SignUp(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	var existingUser User
	if err := db.First(&existingUser, "email = ?", user.Email).Error; err == nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existing email",
		})
		return
	}

	password, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	user.Password = password

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing email or password"})
		return
	}

	var user User
	if err := db.First(&user, "email = ?", loginRequest.Email).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if !CheckHashPassword(loginRequest.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate auth"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth": token})
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
