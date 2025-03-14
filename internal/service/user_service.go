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

type service struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) Login(input request.LoginRequest) (entity.User, error) {
	user, err := s.repository.GetByPhoneNumber(input.PhoneNumber)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) VerifyPin(input request.VerifyPinRequest) (entity.User, error) {
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

func (s *service) GetById(userId int) (entity.User, error) {
	user, err := s.repository.GetById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetByUuid(uuid string) (entity.User, error) {
	user, err := s.repository.GetByUuid(uuid)
	if err != nil {
		return user, err
	}

	return user, nil
}
