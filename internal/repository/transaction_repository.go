package repository

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Store(input entity.Transaction) (entity.Transaction, error)
	Update(input entity.Transaction) error
	GetByConsumerOtp(consumerId uint, otp string) (entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Store(input entity.Transaction) (entity.Transaction, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	return input, nil
}

func (r *transactionRepository) GetByConsumerOtp(consumerId uint, otp string) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.db.Where("consumer_id = ?", consumerId).Where("otp = ?", otp).Find(&transaction).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}

func (r *transactionRepository) Update(input entity.Transaction) error {
	err := r.db.Save(&input).Error
	if err != nil {
		return err
	}

	return nil
}
