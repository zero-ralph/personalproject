package modelsServices

import (
	"errors"
	"recipes/core/models"
)

func GetUserById(userId string) (models.User, error) {
	var user models.User

	if err := models.DBConn.First(&user, "id=?", userId).Error; err != nil {
		return user, errors.New("user not found")
	}

	user.RetractPassword()
	return user, nil
}
