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
	routerGroup.GET("/recipes", controllers.RecipesList)
	routerGroup.GET("/recipes/:id", controllers.RecipeDetails)

}

func PrivateAPI(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middleware.JWTAuthMiddleware())
	routerGroup.Use(middleware.CurrentUser())
	routerGroup.GET("/users", controllers.UsersList)
	routerGroup.POST("/recipes", controllers.CreateRecipe)
	routerGroup.PUT("/recipes/:id", controllers.UpdateRecipe)
	routerGroup.PATCH("/recipes/:id", controllers.PartialUpdateRecipe)
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
