package config

import (
	"github.com/spf13/viper"
)

type ConfigManager struct {
	configInstance *viper.Viper
}

const (
	DEFAULT_CONFIG_FILE_TYPE string = "toml"
	DEFAULT_CONFIG_FILE_NAME string = "config"
)

const (
	app     = "app"
	appPort = app + ".port"

	database         = "database"
	databaseName     = database + ".name"
	databaseHost     = database + ".host"
	databaseUsername = database + ".username"
	databasePassword = database + ".password"
	databasePort     = database + ".port"
	databaseSSLMode  = database + ".sslmode"

	jwt             = "jwt"
	jwtSecret       = jwt + ".secret"
	jwtHourLifeSpan = jwt + ".hourLifeSpan"
)

func NewConfigManager() (configManager *ConfigManager) {
	configManager = &ConfigManager{
		configInstance: viper.New(),
	}
	configManager.configInstance.SetConfigName(DEFAULT_CONFIG_FILE_NAME)
	configManager.configInstance.SetConfigType(DEFAULT_CONFIG_FILE_TYPE)
	return
}
func (configManager *ConfigManager) ReadConfigFile(configFile string) (err error) {
	configManager.configInstance.SetConfigFile(configFile)
	if err = configManager.configInstance.ReadInConfig(); err != nil {
		return
	}
	return
}

func (configManager *ConfigManager) GetDatabaseHost() string {
	return configManager.configInstance.GetString(databaseHost)
}
func (configManager *ConfigManager) GetDatabaseName() string {
	return configManager.configInstance.GetString(databaseName)
}
func (configManager *ConfigManager) GetDatabaseUsername() string {
	return configManager.configInstance.GetString(databaseUsername)
}

func (configManager *ConfigManager) GetDatabasePassword() string {
	return configManager.configInstance.GetString(databasePassword)
}
func (configManager *ConfigManager) GetDatabasePort() int {
	return configManager.configInstance.GetInt(databasePort)
}

func (configManager *ConfigManager) GetDatabaseSSLMode() string {
	return configManager.configInstance.GetString(databaseSSLMode)
}

func (configManager *ConfigManager) GetAppPort() string {
	return configManager.configInstance.GetString(appPort)
}

func (configManager *ConfigManager) GetJWTSecret() string {
	return configManager.configInstance.GetString(jwtSecret)
}

func (configManager *ConfigManager) GetJWTHourLifeSpan() int {
	return configManager.configInstance.GetInt(jwtHourLifeSpan)
}
