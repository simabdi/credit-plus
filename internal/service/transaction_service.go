package service

import (
	"credit-plus/internal/model/entity"
	"credit-plus/internal/repository"
	"credit-plus/internal/request"
	"github.com/gofiber/fiber/v2"
)

type TransactionService interface {
	Store(ctx *fiber.Ctx, input request.TransactionRequest) (entity.Transaction, error)
}

type transactionService struct {
	userRepository     repository.UserRepository
	consumerRepository repository.ConsumerRepository
	limitRepository repository.LimitRepository
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(userRepository repository.UserRepository, consumerRepository repository.ConsumerRepository, limitRepository repository.LimitRepository, transactionRepository repository.TransactionRepository) *transactionService {
	return &transactionService{userRepository, consumerRepository, limitRepository, transactionRepository}
}

func (s *transactionService) Store(ctx *fiber.Ctx, input request.TransactionRequest) (entity.Transaction, error) {
	userLogin, err := s.userRepository.GetByUuid(ctx.Locals("uuid").(string))
	if err != nil {
		return entity.Transaction{}, err
	}

	consumer, err := s.consumerRepository.GetByUserId(userLogin.ID)
	if err != nil {
		return entity.Transaction{}, err
	}

	data := entity.Transaction{
		ConsumerId: consumer.ID,
		ContractNumber: "",
		Otr: ,
		AdminFee: ,
		InstallmentAmount: ,
		AmountOfInterest: ,
		AssetName: input.AssetName,
		Platform: input.Platform,
		Otp: "",
		Status: "Unpaid",
	}

	result, err := s.transactionRepository.Store(data)
	if err != nil {
		return entity.Transaction{}, err
	}

	return result, nil
}