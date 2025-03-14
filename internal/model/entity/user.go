package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey"`
	Uuid        string    `gorm:"type:varchar(100);unique"`
	PhoneNumber string    `gorm:"type:varchar(13);unique"`
	Pin         string    `gorm:"type:char(6)"`
	Status      string    `gorm:"type:varchar(10)"`
	CreatedAt   time.Time `gorm:"<-:create;type:datetime(0)"`
	UpdatedAt   time.Time `gorm:"<-:update;type:datetime(0)"`
	gorm.DeletedAt
}
