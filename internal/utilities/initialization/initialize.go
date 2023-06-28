package initialization

import (
	"fmt"
	"os"
	"recipes/core/models"
	"recipes/core/routes"
	"recipes/utilities/config"
)

func InitializeConfig(settings string) config.ConfigManager {
	configManager := config.NewConfigManager()
	err := configManager.ReadConfigFile(settings)
	if err != nil {
		fmt.Println("Config file error: ", err)
		os.Exit(1)
	}
	return *configManager
}

func InitializeDatabaseConnection(config *config.ConfigManager) {
	models.DatabaseConnection(config)
}

func InitializationExecute(settings string) {
	configManager := InitializeConfig(settings)
	InitializeDatabaseConnection(&configManager)
	fmt.Println("=========================================")
	fmt.Println("=========================================")
	fmt.Println("==         System is Starting          ==")
	fmt.Println("=========================================")
	fmt.Println("=========================================")

	// Routings
	routes.Start(&configManager)
}
