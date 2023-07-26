package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/services"
	"github.com/zero-ralph/personalporject_users/domain_service/pkg/utils"
)

func JWTAuthMiddleware(service *services.DomainServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.ValidateToken(c, *service)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized."})
			return
		}
		c.Next()
	}
}

func UserDomainJWTAuthMiddleware(service *services.UserDomainServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.UserDomainValidateToken(c, *service)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Unauthorized."})
			return
		}
		c.Next()
	}
}
