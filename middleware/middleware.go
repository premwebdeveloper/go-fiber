package middleware

import (
	"web/model/auth"

	"github.com/gofiber/fiber/v2"
)

func ValidateToken(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")

	if tokenString == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Missing Authorization Header")
	}
	tokenString = tokenString[7:]

	username, err := auth.GetTokenForUser(tokenString)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if username == "" {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	ctx.Locals("username", username)

	return ctx.Next()
}
