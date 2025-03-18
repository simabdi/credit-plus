package route

import (
	"credit-plus/internal/config"
	"credit-plus/internal/handler"
	"credit-plus/internal/middleware"
	"credit-plus/internal/repository"
	"credit-plus/internal/service"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) {
	middlewareService := middleware.NewJwtService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, middlewareService)

	limitRepository := repository.NewLimitRepository(db)
	limitService := service.NewLimitService(userRepository, limitRepository)
	limitHandler := handler.NewLimitHandler(limitService)

	consumerRepository := repository.NewConsumerRepository(db)
	parameterRepository := repository.NewParameterRepository(db)

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(userRepository, consumerRepository, limitRepository, transactionRepository, parameterRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app := SetupApp()
	api := app.Group("/api/v1")
	api.Post("/auth/login", userHandler.Login)
	api.Post("/auth/verify-pin", userHandler.VerifyPin)
	authorized := api
	authorized.Use(middleware.Middleware(middlewareService, userService))
	{
		limitGroup := authorized.Group("limits")
		{
			limitGroup.Get("", limitHandler.CheckAllLimit)
			limitGroup.Get("/by-amount", limitHandler.CheckLimitByAmount)
		}

		transactionGroup := authorized.Group("transactions")
		{
			transactionGroup.Post("", transactionHandler.Save)
			transactionGroup.Post("/verify", transactionHandler.Update)
		}
	}

	log.Fatal(app.Listen(":" + config.ListeningPort))
}
