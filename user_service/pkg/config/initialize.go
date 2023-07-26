package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zero-ralph/personalporject_users/user_service/internal/api"
	"github.com/zero-ralph/personalporject_users/user_service/internal/model"
	"github.com/zero-ralph/personalporject_users/user_service/internal/repository"
	"github.com/zero-ralph/personalporject_users/user_service/internal/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Gin    *gin.Engine
	Config *ConfigManager
}

func (s *Server) InitializeDatabase() (server *Server) {
	connection := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
		s.Config.GetDatabaseHost(),
		s.Config.GetDatabasePort(),
		s.Config.GetDatabaseUsername(),
		s.Config.GetDatabaseName(),
		s.Config.GetDatabasePassword(),
		s.Config.GetDatabaseSSLMode())
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	s.DB = db
	s.DB.AutoMigrate(
		&model.User{},
		&model.Profile{},
	)
	return
}

func (s *Server) InitializeRoutes(userService *services.UserServiceInterface) {

	g := gin.Default()
	s.Gin = g
	apiRouter := s.Gin.Group("/api")
	api.NewUserAPI(apiRouter, userService)
}

func (s *Server) Ready() bool {
	return s.DB != nil && s.Gin != nil
}

func (s *Server) ExecuteStart() {
	s.InitializeDatabase()
	userRepository := repository.NewUserRepository(s.DB, s.Config.GetJWTHourLifeSpan(), s.Config.GetJWTSecret())
	userService := services.NewUserService(userRepository)

	s.InitializeRoutes(&userService)

	if !s.Ready() {
		fmt.Println(errors.New("server isn't ready - make sure to init db and gin"))
		os.Exit(1)
	}

	s.Gin.Run(fmt.Sprintf(":%v", s.Config.GetAppPort()))

}
