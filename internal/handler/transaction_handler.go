package handler

import (
	"credit-plus/internal/exception"
	"credit-plus/internal/helper"
	"credit-plus/internal/request"
	"credit-plus/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type transactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) Save(ctx *fiber.Ctx) error {
	var input request.TransactionRequest

	if err := ctx.BodyParser(&input); err != nil {
		JsonResponse := helper.JsonResponse(http.StatusUnprocessableEntity, "", false, exception.Validation(input), nil)
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(JsonResponse)
	}

	save, err := h.transactionService.Store(ctx, input)
	if err != nil {
		JsonResponse := helper.JsonResponse(fiber.StatusBadRequest, "", false, exception.Error(err), nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(JsonResponse)
	}

	JsonResponse := helper.JsonResponse(http.StatusOK, "", true, "", (save))
	return ctx.Status(fiber.StatusOK).JSON(JsonResponse)
}
