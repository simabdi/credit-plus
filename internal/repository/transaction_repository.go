package repository

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Store(input entity.Transaction) (entity.Transaction, error)
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
