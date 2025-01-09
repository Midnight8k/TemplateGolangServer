package repository

import (
	"blankproject/entities"

	"gorm.io/gorm"
)

type MessageRepository interface {
	Save(message *entities.Message) error
	GetAll() ([]entities.Message, error)
}

type GormMessageRepository struct {
	DB *gorm.DB
}

func (r *GormMessageRepository) Save(message *entities.Message) error {
	return r.DB.Create(message).Error
}

func (r *GormMessageRepository) GetAll() ([]entities.Message, error) {
	var messages []entities.Message
	if err := r.DB.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
