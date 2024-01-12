package service

import (
	"github.com/felipecveiga/crud_estudo_teste/model"
	"github.com/felipecveiga/crud_estudo_teste/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(useRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: useRepo,
	}
}

func (s *UserService) GetAllUsers() ([]model.User,error) {
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}