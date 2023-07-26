package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/middleware"
	"github.com/zero-ralph/personalporject_users/domain_service/internal/services"
)

type UserDomainAPIInterface interface {
	GetUserDomains(c *gin.Context)
}

type UserDomainAPIStruct struct {
	service services.UserDomainServiceInterface
}

func NewUserDomainAPI(api *gin.RouterGroup, service *services.UserDomainServiceInterface) {
	controller := &UserDomainAPIStruct{
		service: *service,
	}

	api.Use(middleware.UserDomainJWTAuthMiddleware(service))
	api.GET("/domain/get-by/:userId", controller.GetUserDomains)
}

func (userDomainAPI *UserDomainAPIStruct) GetUserDomains(c *gin.Context) {
	userId, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	domains, err := userDomainAPI.service.GetUserDomains(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": domains})
}
