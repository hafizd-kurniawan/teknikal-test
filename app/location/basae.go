package location

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	location := router.Group("/locations")
	location.Post("/", handler.CreateLocation)
	location.Get("/", handler.GetAllLocations)
	location.Get("/:id", handler.GetLocationByID)
	location.Put("/:id", handler.UpdateLocation)
	location.Delete("/:id", handler.DeleteLocation)
}
