package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDomainPivot struct {
	ID       uuid.UUID `gorm:"type:uuid; primaryKey" json:"id"`
	DomainId uuid.UUID `json:"domain_id"`
	UserId   uuid.UUID `gorm:"type:uuid; not null;" json:"user_id"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func (userDomain *UserDomainPivot) BeforeCreate(tx *gorm.DB) (err error) {
	userDomain.ID = uuid.New()
	return nil
}
