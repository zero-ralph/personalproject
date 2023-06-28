package routes

import (
	"fmt"
	"recipes/core/controllers"
	"recipes/core/middleware"
	"recipes/utilities/config"

	"github.com/gin-gonic/gin"
)

func PublicAPI(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/login", controllers.Login)
	routerGroup.POST("/register", controllers.Register)

}

func PrivateAPI(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middleware.JWTAuthMiddleware())
	routerGroup.Use(middleware.CurrentUser())
	routerGroup.GET("/users", controllers.UsersList)
}

func Start(config *config.ConfigManager) {

	if !config.GetAppDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middleware.Congfiguration(config))

	routerGroup := router.Group("/api")
	PublicAPI(routerGroup)
	PrivateAPI(routerGroup)

	router.Run(fmt.Sprintf(":%v", config.GetAppPort()))
}
