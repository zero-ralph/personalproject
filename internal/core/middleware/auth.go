package middleware

import (
	"net/http"
	"recipes/utilities/config"
	"recipes/utilities/token"

	"github.com/gin-gonic/gin"
)

func Congfiguration(config *config.ConfigManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized.")
			c.Abort()
			return
		}
		c.Next()
	}
}
