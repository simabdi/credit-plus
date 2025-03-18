package entity

import "time"

type Transaction struct {
	ID                uint `gorm:"primaryKey"`
	ConsumerId        uint
	ContractNumber    string    `gorm:"type:varchar(30)"`
	Otr               int       `gorm:"type:integer(11)"`
	AdminFee          int       `gorm:"type:integer(11)"`
	InstallmentAmount int       `gorm:"type:tinyint(2)"`
	AmountOfInterest  float32   `gorm:"type:float(4)"`
	AssetName         string    `gorm:"type:varchar(80)"`
	Platform          string    `gorm:"type:varchar(50)"`
	Otp               string    `gorm:"type:char(6)"`
	Status            string    `gorm:"type:varchar(10)"`
	CreatedAt         time.Time `gorm:"<-:create;type:datetime(0)"`
	UpdatedAt         time.Time `gorm:"<-:update;type:datetime(0)"`
}
