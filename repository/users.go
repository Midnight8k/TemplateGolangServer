package repository

import (
	"blankproject/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entities.User) error
}

type GormUserRepository struct {
	DB *gorm.DB
}

func (r *GormUserRepository) Save(user *entities.User) error {
	return r.DB.Create(user).Error
}
