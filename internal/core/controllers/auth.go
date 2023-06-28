package controllers

import (
	"net/http"
	"recipes/core/models"
	"recipes/utilities/config"
	"recipes/utilities/token"

	"github.com/gin-gonic/gin"
)

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input AuthenticationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	user.Save()

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Register Successful",
		"user":    user,
	})
}

func Login(c *gin.Context) {

	var input AuthenticationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.ValidateCredentials(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Password is incorrect."})
	}

	config := c.MustGet("config").(*config.ConfigManager)
	token, err := token.GenerateToken(user.ID, config)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't generate token."})
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})

}
