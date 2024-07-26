package location

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"employee-management/infra/middleware"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	location := router.Group("/locations")
	location.Post("/", middleware.AuthMiddleware(), handler.CreateLocation)
	location.Get("/", middleware.AuthMiddleware(), handler.GetAllLocations)
	location.Get("/:id", middleware.AuthMiddleware(), handler.GetLocationByID)
	location.Put("/:id", middleware.AuthMiddleware(), handler.UpdateLocation)
	location.Delete("/:id", middleware.AuthMiddleware(), handler.DeleteLocation)
}
