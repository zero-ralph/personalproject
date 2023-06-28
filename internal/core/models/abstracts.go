package models

import (
	"time"

	"github.com/google/uuid"
)

type UUIDModel struct {
	ID uuid.UUID `gorm:"type:uuid; primary_key" json:"id"`
}

type TimeStampedModel struct {
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type SoftDeleteModel struct {
	IsRemoved bool `gorm:"default:false" json:"is_removed"`
}

type ActiveInactive struct {
	IsActive bool `gorm:"default:false" json:"is_active"`
}
