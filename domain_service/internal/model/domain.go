package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Domain struct {
	ID      uuid.UUID `gorm:"type:uuid; primaryKey" json:"id"`
	Name    string    `gorm:"size:100; not null; unique" json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (domain *Domain) BeforeCreate(tx *gorm.DB) (err error) {
	domain.ID = uuid.New()
	return nil
}
