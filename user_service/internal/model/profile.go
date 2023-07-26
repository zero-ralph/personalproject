package model

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID         uuid.UUID `gorm:"type:uuid; primarykey" json:"id"`
	Firstname  string    `gorm:"size:100; not null;" json:"firstname"`
	Middlename string    `gorm:"size:100;" json:"middlename"`
	Lastname   string    `gorm:"size:100; not null;" json:"lastname"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	UserId     uuid.UUID `json:"user_id"`
}
