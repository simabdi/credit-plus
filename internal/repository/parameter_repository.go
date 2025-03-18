package repository

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
)

type ParameterRepository interface {
	GetWhere(column, value string) (entity.Parameter, error)
}

type parameterRepository struct {
	db *gorm.DB
}

func NewParameterRepository(db *gorm.DB) *parameterRepository {
	return &parameterRepository{db}
}

func (r *parameterRepository) GetWhere(column, value string) (entity.Parameter, error) {
	var parameter entity.Parameter
	err := r.db.Where(column+" = ?", value).First(&parameter).Error
	if err != nil {
		return entity.Parameter{}, err
	}

	return parameter, nil
}
