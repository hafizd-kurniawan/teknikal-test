package attendacne

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	attendance := router.Group("/attendances")
	attendance.Post("/", handler.CreateAttendance)
	attendance.Get("/", handler.GetAllAttendances)
	attendance.Get("/:id", handler.GetAttendanceByID)
	attendance.Put("/:id", handler.UpdateAttendance)
	attendance.Delete("/:id", handler.DeleteAttendance)
}
