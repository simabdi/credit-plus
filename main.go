package main

import (
	"credit-plus/internal/config"
	"credit-plus/internal/database"
	"credit-plus/route"
)

func main() {
	config.Initialize()
	db := config.Connection()
	database.Migrate(db)
	route.Initialize(db)
}
