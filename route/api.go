package route

import (
	"credit-plus/internal/config"
	"credit-plus/internal/handler"
	"credit-plus/internal/middleware"
	"credit-plus/internal/repository"
	"credit-plus/internal/service"
	log "github.com/sirupsen/logrus"
)

func Initialize() {
	db := config.Connection()

	middlewareService := middleware.NewJwtService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, middlewareService)

	app := SetupApp()
	api := app.Group("/api/v1")
	api.Post("/auth/login", userHandler.Login)
	authorized := api
	authorized.Use(middleware.Middleware(middlewareService, userService))
	{

	}

	log.Fatal(app.Listen(":" + config.ListeningPort))
}
