package service

import (
	"errors"

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

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserID(userID string) (*model.User, error) {
	user, err := s.UserRepository.GetUserFromID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) RemoveUserByID(userID string) (*model.User, error) {
	user, err := s.UserRepository.RemoveUserFromID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CreateUser(user *model.User) error {
	existingUser, err := s.UserRepository.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("Email já está em uso")
	}
	return s.UserRepository.CreateUserFromBD(user)
}
