package repository

import (
	"github.com/felipecveiga/crud_estudo_teste/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: *db,
	}
}

func (r UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}