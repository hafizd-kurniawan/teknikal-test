package employee

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	_ = handler

	{
		employee := router.Group("/employees")
		employee.Post("/", handler.CreateEmployee)
		employee.Get("/", handler.GetAllEmployees)
		employee.Get("/:id", handler.GetEmployeeByID)
		employee.Put("/:id", handler.UpdateEmployee)
		employee.Delete("/:id", handler.DeleteEmployee)
	}
}
