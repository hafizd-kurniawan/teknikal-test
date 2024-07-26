package employee

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"employee-management/infra/middleware"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	_ = handler

	{
		employee := router.Group("/employees")
		employee.Post("/login", handler.Login)
		employee.Post("/", handler.CreateEmployee)
		employee.Get("/", middleware.AuthMiddleware(), handler.GetAllEmployees)
		employee.Get("/:id", middleware.AuthMiddleware(), handler.GetEmployeeByID)
		employee.Put("/:id", middleware.AuthMiddleware(), handler.UpdateEmployee)
		employee.Delete("/:id", middleware.AuthMiddleware(), handler.DeleteEmployee)

	}
}
