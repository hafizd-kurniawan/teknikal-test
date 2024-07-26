package attendacne

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"employee-management/infra/middleware"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	attendance := router.Group("/attendances")
	attendance.Post("/", middleware.AuthMiddleware(), handler.CreateAttendance)
	attendance.Get("/", middleware.AuthMiddleware(), handler.GetAllAttendances)
	attendance.Get("/:id", middleware.AuthMiddleware(), handler.GetAttendanceByID)
	attendance.Put("/:id", middleware.AuthMiddleware(), handler.UpdateAttendance)
	attendance.Put("/", middleware.AuthMiddleware(), handler.UpdateCheckoutAttendance)
	attendance.Delete("/:id", middleware.AuthMiddleware(), handler.DeleteAttendance)
}
