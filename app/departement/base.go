package departement

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"employee-management/infra/middleware"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	department := router.Group("/departments")
	{

		department.Post("/",middleware.AuthMiddleware(), handler.CreateDepartment)
		department.Get("/",middleware.AuthMiddleware(), handler.GetAllDepartments)
		department.Get("/:id",middleware.AuthMiddleware(), handler.GetDepartmentByID)
		department.Put("/:id",middleware.AuthMiddleware(), handler.UpdateDepartment)
		department.Delete("/:id", middleware.AuthMiddleware(), handler.DeleteDepartment)
	}

}
