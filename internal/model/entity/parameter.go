package entity

type Parameter struct {
	ID            uint   `gorm:"primaryKey"`
	ParameterType string `gorm:"type:varchar(20);unique"`
	Value         string `gorm:"type:varchar(15)"`
}
