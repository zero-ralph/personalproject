package models

import (
	"fmt"
	"recipes/utilities/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var err error
var DBConn *gorm.DB

func DatabaseConnection(config *config.ConfigManager) {

	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
		config.GetDatabaseHost(),
		config.GetDatabasePort(),
		config.GetDatabaseUsername(),
		config.GetDatabaseName(),
		config.GetDatabasePassword(),
		config.GetDatabaseSSLMode())

	DBConn, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
	}

	DBConn.AutoMigrate(
		&User{},
		&Recipe{},
	)
}

func VerifyPassword(password string, hashedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidateCredentials(username string, password string) (User, error) {
	u := User{}

	if err := DBConn.Model(User{}).Where("username=?", username).Take(&u).Error; err != nil {
		fmt.Println(err)
	}

	err = VerifyPassword(password, u.Password)
	if err != nil {
		fmt.Println(err)
	}

	return u, err
}
