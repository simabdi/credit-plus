package entity

import "time"

type Consumer struct {
	ID          uint   `gorm:"primaryKey"`
	Uuid        string `gorm:"type:varchar(100);unique"`
	UserId      uint
	Nik         string `gorm:"type:varchar(16);unique"`
	FullName    string `gorm:"type:varchar(50)"`
	LegalName   string `gorm:"type:varchar(50)"`
	PlaceDob    string `gorm:"type:varchar(50)"`
	Dob         string `gorm:"type:date"`
	Salary      string `gorm:"type:varchar(50)"`
	KtpImage    string `gorm:"type:varchar(50)"`
	SelfieImage string `gorm:"type:varchar(50)"`
	Limit       []Limit
	CreatedAt   time.Time `gorm:"<-:create;type:datetime(0)"`
	UpdatedAt   time.Time `gorm:"<-:update;type:datetime(0)"`
}
