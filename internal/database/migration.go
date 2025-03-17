package database

import (
	"credit-plus/internal/model/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	status := db.Migrator().HasTable(&entity.User{})
	if status == false {
		db.AutoMigrate(&entity.User{})
	}

	status = db.Migrator().HasTable(&entity.Consumer{})
	if status == false {
		db.AutoMigrate(&entity.Consumer{})
	}

	status = db.Migrator().HasTable(&entity.Limit{})
	if status == false {
		db.AutoMigrate(&entity.Limit{})
	}

	status = db.Migrator().HasTable(&entity.Transaction{})
	if status == false {
		db.AutoMigrate(&entity.Transaction{})
	}
}
