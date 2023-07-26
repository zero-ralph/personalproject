package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID         uuid.UUID `gorm:"type:uuid; primarykey" json:"id"`
	Firstname  string    `gorm:"size:100; not null;" json:"firstname"`
	Middlename string    `gorm:"size:100;" json:"middlename"`
	Lastname   string    `gorm:"size:100; not null;" json:"lastname"`
	Created    time.Time
	Updated    time.Time
	UserId     uuid.UUID
}

func (profile *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	profile.ID = uuid.New()
	return nil
}
