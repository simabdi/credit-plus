package config

import (
	"credit-plus/internal/exception"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUsername, DbPassword, DbHost, DbPort, DbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"dsn": dsn,
			"db":  db,
			"err": err,
		}).Warning("Connection")

		panic(exception.Error(err))
	}

	log.WithFields(log.Fields{
		"db": "✅ Database connected successfully",
	}).Info("Connection")

	return db
}
