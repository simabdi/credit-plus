package entity

import "time"

type Limit struct {
	ID         uint      `gorm:"primaryKey"`
	Uuid       string    `gorm:"type:varchar(100);unique"`
	ConsumerId uint      `gorm:"index;not null"`
	Tenor      int       `gorm:"type:tinyint(2)"`
	Amount     int       `gorm:"type:integer(11)"`
	CreatedAt  time.Time `gorm:"<-:create;type:datetime(0)"`
	UpdatedAt  time.Time `gorm:"<-:update;type:datetime(0)"`
}
