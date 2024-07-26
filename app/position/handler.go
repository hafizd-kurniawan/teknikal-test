package position

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service
}

func newHandler(service service) handler {
	return handler{service: service}
}
func (h *handler) CreatePosition(c *fiber.Ctx) error {
	req := CreatePositionRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.CreatedBy = "system" // replace with actual user
	// req.CreatedAt = time.Now()

	if err := h.service.CreatePosition(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Position created successfully"})
}

func (h *handler) GetAllPositions(c *fiber.Ctx) error {
	positions, err := h.service.GetAllPositions(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(positions)
}

func (h *handler) GetPositionByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}
	position, err := h.service.GetPositionByID(context.Background(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(position)
}

func (h *handler) UpdatePosition(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}

	req := UpdatePositionRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.PositionID = id
	req.UpdatedBy = "system" // replace with actual user
	// req.UpdatedAt = time.Now()

	if err := h.service.UpdatePosition(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Position updated successfully"})
}

func (h *handler) DeletePosition(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid position ID"})
	}

	if err := h.service.DeletePosition(context.Background(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Position deleted successfully"})
}
