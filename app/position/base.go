package position

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"employee-management/infra/middleware"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(&repo)
	handler := newHandler(svc)

	position := router.Group("/positions")
	{
		position.Post("/",middleware.AuthMiddleware(), handler.CreatePosition)
		position.Get("/",middleware.AuthMiddleware(), handler.GetAllPositions)
		position.Get("/:id",middleware.AuthMiddleware(), handler.GetPositionByID)
		position.Put("/:id",middleware.AuthMiddleware(), handler.UpdatePosition)
		position.Delete("/:id",middleware.AuthMiddleware(), handler.DeletePosition)
	}

}
