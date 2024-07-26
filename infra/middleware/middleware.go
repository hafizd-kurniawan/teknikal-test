package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenHeader := ctx.Get("Authorization")
		if tokenHeader == "" {
			return WriteError(ctx, ErrUnAuthorized)
		}

		claims, err := GetJWTClaims(tokenHeader)
		if err != nil {
			return WriteError(ctx, err)
		}

		id := claims.EmployeeId
		employeeCode := claims.EmployeeCode

		ctx.Locals("id", id)
		ctx.Locals("email", employeeCode)

		return ctx.Next()
	}
}
