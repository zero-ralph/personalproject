package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid; primarykey" json:"id"`
	Username string    `gorm:"size:100;not null;unique" json:"username"`
	Password string    `gorm:"size:100;not null;" json:"password"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Profile  Profile   `json:"profile"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("cannot generate password")
	}
	user.Password = string(password)

	return nil
}

func (user *User) VerifyPassword(password string, hashedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
