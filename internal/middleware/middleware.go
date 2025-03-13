package middleware

import (
	"credit-plus/internal/helper"
	"credit-plus/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func Middleware(service Service, userService service.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		headers := ctx.GetReqHeaders()["Authorization"]
		header := strings.Join(headers, " ")

		if !strings.Contains(header, "Bearer") {
			response := helper.JsonResponse(http.StatusUnauthorized, "Unauthorized", false, "", nil)
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)
		}

		token := ""
		authorization := strings.Split(header, " ")
		if len(authorization) == 2 {
			token = authorization[1]
		}

		validateToken, err := service.VerifyToken(token)
		if err != nil || !validateToken.Valid {
			response := helper.JsonResponse(http.StatusUnauthorized, "Unauthorized", false, err.Error(), err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)
		}

		claim, ok := validateToken.Claims.(jwt.MapClaims)
		if !ok {
			response := helper.JsonResponse(http.StatusUnauthorized, "Unauthorized", false, "", err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)
		}

		_, err = userService.GetByUuid(claim["uuid"].(string))
		if err != nil {
			response := helper.JsonResponse(http.StatusUnauthorized, "Unauthorized", false, err.Error(), err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(response)
		}

		ctx.Locals("uuid", claim["uuid"].(string))
		return ctx.Next()
	}
}
