package departement

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	department := router.Group("/departments")
	{

		department.Post("/", handler.CreateDepartment)
		department.Get("/", handler.GetAllDepartments)
		department.Get("/:id", handler.GetDepartmentByID)
		department.Put("/:id", handler.UpdateDepartment)
		department.Delete("/:id", handler.DeleteDepartment)
	}

}
