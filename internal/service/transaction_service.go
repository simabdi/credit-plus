package service

import (
	"credit-plus/internal/helper"
	"credit-plus/internal/model/entity"
	"credit-plus/internal/repository"
	"credit-plus/internal/request"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

type TransactionService interface {
	Store(ctx *fiber.Ctx, input request.TransactionRequest) (entity.Transaction, error)
}

type transactionService struct {
	userRepository        repository.UserRepository
	consumerRepository    repository.ConsumerRepository
	limitRepository       repository.LimitRepository
	transactionRepository repository.TransactionRepository
	parameterRepository   repository.ParameterRepository
}

func NewTransactionService(userRepository repository.UserRepository, consumerRepository repository.ConsumerRepository, limitRepository repository.LimitRepository, transactionRepository repository.TransactionRepository, parameterRepository repository.ParameterRepository) *transactionService {
	return &transactionService{userRepository, consumerRepository, limitRepository, transactionRepository, parameterRepository}
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

	limit, err := s.limitRepository.GetByUuid(input.UuidLimit)
	if err != nil {
		return entity.Transaction{}, err
	}

	if limit.CurrentAmount < input.Otr {
		return entity.Transaction{}, errors.New("Limit not enough")
	}

	parameterAdminFee, err := s.parameterRepository.GetWhere("parameter_type", "admin_fee")
	if err != nil {
		return entity.Transaction{}, err
	}
	adminFee, _ := strconv.Atoi(parameterAdminFee.Value)
	ParameterInterest, err := s.parameterRepository.GetWhere("parameter_type", "interest")
	if err != nil {
		return entity.Transaction{}, err
	}
	interest, _ := strconv.ParseFloat(strings.TrimSpace(ParameterInterest.Value), 64)
	AmountOfInterest := ((interest * float64(input.Otr)) / 100) * float64(limit.Tenor)

	data := entity.Transaction{
		ConsumerId:        consumer.ID,
		ContractNumber:    helper.InvoiceNumber(),
		Otr:               input.Otr,
		AdminFee:          adminFee,
		InstallmentAmount: limit.Tenor,
		AmountOfInterest:  float32(AmountOfInterest),
		AssetName:         input.AssetName,
		Platform:          input.Platform,
		Otp:               helper.GenerateOtp(),
		Status:            "Unpaid",
	}

	result, err := s.transactionRepository.Store(data)
	if err != nil {
		return entity.Transaction{}, err
	}

	/**
	 * Send OTP via phone number
	 */
	//Here send otp

	return result, nil
}
