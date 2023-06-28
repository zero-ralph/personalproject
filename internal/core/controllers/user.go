package controllers

import (
	"net/http"
	"recipes/core/models"

	"github.com/gin-gonic/gin"
)

func UsersList(c *gin.Context) {
	var users []models.User
	models.DBConn.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"count": len(users),
		"data":  users,
	})
}
