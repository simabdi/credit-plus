package config

import (
	"credit-plus/internal/exception"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

var DbHost string
var DbPort string
var DbDatabase string
var DbUsername string
var DbPassword string
var ListeningPort string
var JWTSecretKey string
var LifeTime string
var UrlImage string

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warning("Cannot load env file")

		panic(exception.Error(err))
	}

	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbDatabase = os.Getenv("DB_DATABASE")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	ListeningPort = os.Getenv("PORT")
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	LifeTime = os.Getenv("LIFETIME")
	UrlImage = os.Getenv("URL_IMAGE")
}
