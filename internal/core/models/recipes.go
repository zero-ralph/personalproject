package models

import (
	"errors"

	"github.com/google/uuid"
)

type Recipe struct {
	UUIDModel
	Title   string `gorm:"size:100;not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	TimeStampedModel
	SoftDeleteModel
	ActiveInactive
	Author uuid.UUID `gorm:"type:uuid;column:author_id;not null;constraint:OnDelete:CASCADE;" json:"author"`
}

func (recipe *Recipe) Save() (*Recipe, error) {

	recipe.UUIDModel.ID = uuid.New()
	err := DBConn.Create(&recipe).Error

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return recipe, nil
}
