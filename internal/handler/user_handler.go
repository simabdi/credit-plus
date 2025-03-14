package handler

import (
	"credit-plus/internal/exception"
	"credit-plus/internal/helper"
	"credit-plus/internal/middleware"
	resource "credit-plus/internal/model/resources"
	"credit-plus/internal/request"
	"credit-plus/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type userHandler struct {
	service           service.UserService
	middlewareService middleware.Service
}

func NewUserHandler(service service.UserService, middlewareService middleware.Service) *userHandler {
	return &userHandler{service, middlewareService}
}

func (h *userHandler) Login(ctx *fiber.Ctx) error {
	var input request.LoginRequest

	if err := ctx.BodyParser(&input); err != nil {
		JsonResponse := helper.JsonResponse(http.StatusUnprocessableEntity, "", false, exception.Error(err), nil)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(JsonResponse)
	}

	userLogin, err := h.service.Login(input)
	if err != nil {
		JsonResponse := helper.JsonResponse(http.StatusBadRequest, "phone number is not association. Please sign up for login", false, "", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(JsonResponse)
	}

	JsonResponse := helper.JsonResponse(http.StatusOK, "", true, "", resource.CheckAccountResource(userLogin))
	return ctx.Status(fiber.StatusOK).JSON(JsonResponse)
}

func (h *userHandler) VerifyPin(ctx *fiber.Ctx) error {
	var input request.VerifyPinRequest

	if err := ctx.BodyParser(&input); err != nil {
		JsonResponse := helper.JsonResponse(http.StatusBadRequest, "", false, exception.Error(err), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(JsonResponse)
	}

	userLogin, err := h.service.VerifyPin(input)
	if err != nil {
		JsonResponse := helper.JsonResponse(http.StatusBadRequest, "Pin incorrect.", false, "", nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(JsonResponse)
	}

	err = exception.Validate.Struct(input)
	if err != nil {
		JsonResponse := helper.JsonResponse(http.StatusUnprocessableEntity, "", false, exception.Validation(input), nil)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(JsonResponse)
	}

	token, err := h.middlewareService.GenerateToken(userLogin)
	if err != nil {
		JsonResponse := helper.JsonResponse(http.StatusBadRequest, "", false, exception.Error(err), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(JsonResponse)
	}

	JsonResponse := helper.JsonResponse(http.StatusOK, "", true, "", resource.LoginResource(userLogin, token))
	return ctx.Status(fiber.StatusOK).JSON(JsonResponse)
}
