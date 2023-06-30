package middleware

import (
	"net/http"
	modelsServices "recipes/core/modelServices"
	"recipes/utilities/token"

	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := token.ExtractTokenID(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user, err := modelsServices.GetUserById(userId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.Set("currentUser", user)
		c.Next()
	}

}
