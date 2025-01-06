package repository

import (
	"blankproject/entities"

	"gorm.io/gorm"
)

type MessageRepository interface {
	Save(message *entities.Message) error
}

type GormMessageRepository struct {
	DB *gorm.DB
}

func (r *GormMessageRepository) Save(message *entities.Message) error {
	return r.DB.Create(message).Error
}
