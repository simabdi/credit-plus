package service

import (
	"credit-plus/internal/model/entity"
	"credit-plus/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type LimitService interface {
	GetAll(ctx *fiber.Ctx) ([]entity.Limit, error)
	GetByAmount(amount int) ([]entity.Limit, error)
}

type limitService struct {
	userRepository  repository.UserRepository
	limitRepository repository.LimitRepository
}

func NewLimitService(userRepository repository.UserRepository, limitRepository repository.LimitRepository) *limitService {
	return &limitService{userRepository, limitRepository}
}

func (s *limitService) GetAll(ctx *fiber.Ctx) ([]entity.Limit, error) {
	userUuid := ctx.Locals("uuid").(string)
	userLogin, err := s.userRepository.GetByUuid(userUuid)
	if err != nil {
		return nil, nil
	}

	result, err := s.limitRepository.GetAll(userLogin.ID)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func (s *limitService) GetByAmount(amount int) ([]entity.Limit, error) {
	result, err := s.limitRepository.GetByAmount(amount)
	if err != nil {
		return nil, nil
	}

	return result, nil
}
