package models

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUIDModel
	Username string `gorm:"size:100;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null;unique" json:"password"`
	TimeStampedModel
	SoftDeleteModel
	ActiveInactive
}

func (user *User) HashPassword() *User {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Cannot Generate Password")
		os.Exit(1)
	}
	user.Password = string(hashPassword)
	return user
}

func (user *User) Save() *User {

	user.UUIDModel.ID = uuid.New()
	user.HashPassword()
	err := DBConn.Create(&user).Error

	if err != nil {
		return &User{}
	}

	return user
}

func (user *User) VerifyPassword(passwordInput string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordInput))
}

func (user *User) RetractPassword() {
	user.Password = ""
}
