package repository

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
)

type LimitRepository interface {
	GetAll(userId uint) ([]entity.Limit, error)
	GetByUuid(uuid string) (entity.Limit, error)
	GetByAmount(amount int) ([]entity.Limit, error)
}

type limitRepository struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) *limitRepository {
	return &limitRepository{db}
}

func (r *limitRepository) GetAll(userId uint) ([]entity.Limit, error) {
	var limit []entity.Limit
	err := r.db.Where("user_id = ?", userId).Find(&limit).Error
	if err != nil {
		return nil, err
	}

	return limit, nil
}

func (r *limitRepository) GetByUuid(uuid string) (entity.Limit, error) {
	var limit entity.Limit
	err := r.db.Where("uuid = ?", uuid).Find(&limit).Error
	if err != nil {
		return entity.Limit{}, err
	}

	return limit, nil
}

func (r *limitRepository) GetByAmount(amount int) ([]entity.Limit, error) {
	var limit []entity.Limit
	err := r.db.Where("current_amount >= ?", amount).Find(&limit).Error
	if err != nil {
		return nil, err
	}

	return limit, nil
}
