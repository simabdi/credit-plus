package repository

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
)

type ConsumerRepository interface {
	GetByUserId(userId uint) (entity.Consumer, error)
}

type consumerRepository struct {
	db *gorm.DB
}

func NewConsumerRepository(db *gorm.DB) *consumerRepository {
	return &consumerRepository{db}
}

func (r *consumerRepository) GetByUserId(userId uint) (entity.Consumer, error) {
	var consumer entity.Consumer
	err := r.db.Where("user_id = ?", userId).First(&consumer).Error
	if err != nil {
		return entity.Consumer{}, err
	}

	return consumer, nil
}
