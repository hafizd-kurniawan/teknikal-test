package middleware

import (
	"fmt"

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

		id := claims.EmployeeID
		employeeCode := claims.EmployeeCode
		employeeName := claims.EmployeeName

		fmt.Println(employeeName)
		ctx.Locals("employee_id", id)
		ctx.Locals("employee_code", employeeCode)
		ctx.Locals("employee_name", employeeName)

		return ctx.Next()
	}
}
