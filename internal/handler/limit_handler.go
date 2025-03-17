package handler

import (
	"credit-plus/internal/exception"
	"credit-plus/internal/helper"
	resource "credit-plus/internal/model/resources"
	"credit-plus/internal/request"
	"credit-plus/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type limitHandler struct {
	limitService service.LimitService
}

func NewLimitHandler(limitService service.LimitService) *limitHandler {
	return &limitHandler{limitService}
}

func (h *limitHandler) CheckAllLimit(ctx *fiber.Ctx) error {
	limit, err := h.limitService.GetAll(ctx)
	if err != nil {
		JsonResponse := helper.JsonResponse(http.StatusBadRequest, "", false, exception.Error(err), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(JsonResponse)
	}

	JsonResponse := helper.JsonResponse(http.StatusOK, "", true, "", resource.LimitCollectionResource(limit))
	return ctx.Status(fiber.StatusOK).JSON(JsonResponse)
}

func (h *limitHandler) CheckLimitByAmount(ctx *fiber.Ctx) error {
	var input request.LimitRequest

	if err := ctx.BodyParser(&input); err != nil {
		JsonResponse := helper.JsonResponse(http.StatusBadRequest, "", false, exception.Error(err), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(JsonResponse)
	}

	err := exception.Validate.Struct(input)
	if err != nil {
		JsonResponse := helper.JsonResponse(http.StatusUnprocessableEntity, "", false, exception.Validation(input), nil)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(JsonResponse)
	}

	limit, err := h.limitService.GetByAmount(input.Amount)

	JsonResponse := helper.JsonResponse(http.StatusOK, "", true, "", resource.LimitCollectionResource(limit))
	return ctx.Status(fiber.StatusOK).JSON(JsonResponse)
}
