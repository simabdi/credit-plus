package repository

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserRepository interface {
	Store(input entity.User) (entity.User, error)
	GetByPhoneNumber(phoneNumber string) (entity.User, error)
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

func (r *repository) GetByPhoneNumber(phoneNumber string) (entity.User, error) {
	var user entity.User

	err := r.db.Scopes(statusActive).Where("phone_number = ?", phoneNumber).First(&user).Error
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

func (r *repository) Store(input entity.User) (entity.User, error) {
	err := r.db.Create(&input).Error
	if err != nil {
		return input, err
	}

	return input, nil
}
