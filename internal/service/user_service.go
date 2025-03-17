package service

import (
	"credit-plus/internal/model/entity"
	"credit-plus/internal/repository"
	"credit-plus/internal/request"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(input request.LoginRequest) (entity.User, error)
	VerifyPin(input request.VerifyPinRequest) (entity.User, error)
	GetById(userId int) (entity.User, error)
	GetByUuid(uuid string) (entity.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Login(input request.LoginRequest) (entity.User, error) {
	user, err := s.repository.GetByPhoneNumber(input.PhoneNumber)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) VerifyPin(input request.VerifyPinRequest) (entity.User, error) {
	user, err := s.repository.GetByUuid(input.Uuid)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(input.Pin))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) GetById(userId int) (entity.User, error) {
	user, err := s.repository.GetById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) GetByUuid(uuid string) (entity.User, error) {
	user, err := s.repository.GetByUuid(uuid)
	if err != nil {
		return user, err
	}

	return user, nil
}
