package model

import (
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

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err == nil {
		tx.Statement.SetColumn("password", string(password))
	}

	tx.Statement.SetColumn("Updated", time.Now())
	return

}

func (user *User) VerifyOldPassword(oldPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
}
