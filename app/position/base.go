package position

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	position := router.Group("/positions")
	{
		position.Post("/", handler.CreatePosition)
		position.Get("/", handler.GetAllPositions)
		position.Get("/:id", handler.GetPositionByID)
		position.Put("/:id", handler.UpdatePosition)
		position.Delete("/:id", handler.DeletePosition)
	}

}
