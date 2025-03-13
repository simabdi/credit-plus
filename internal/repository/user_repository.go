package repository

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserRepository interface {
	GetByEmail(email string) (entity.User, error)
	GetById(userId int) (entity.User, error)
	GetByUuid(uuid string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func statusActive(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", "active")
}

func (r *repository) GetByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Scopes(statusActive).Where("email = ?", email).First(&user).Error
	r.db.Logger = logger.Default.LogMode(logger.Info)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetById(userId int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetByUuid(uuid string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
