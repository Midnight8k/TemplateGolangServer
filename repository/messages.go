package repository

import (
	"blankproject/entities"

	"gorm.io/gorm"
)

type MessageRepository interface {
	Save(message *entities.Message) error
	GetAll() ([]entities.Message, error)
	GetById(id int64) (*entities.Message, error)
}

type GormMessageRepository struct {
	DB *gorm.DB
}

func (r *GormMessageRepository) Update(message *entities.Message) error {
	return r.DB.Save(message).Error
}

func (r *GormMessageRepository) Save(message *entities.Message) error {
	return r.DB.Create(message).Error
}

func (r *GormMessageRepository) GetById(id int64) (*entities.Message, error) {
	var message entities.Message
	if err := r.DB.First(&message, id).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *GormMessageRepository) GetAll() ([]entities.Message, error) {
	var messages []entities.Message
	if err := r.DB.Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
