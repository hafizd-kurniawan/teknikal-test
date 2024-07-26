package attendacne

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
func (h *handler) CreateAttendance(c *fiber.Ctx) error {
	req := CreateAttendanceRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.CreatedBy = "system" // replace with actual user
	// req.CreatedAt = time.Now()

	if err := h.service.CreateAttendance(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Attendance created successfully"})
}

func (h *handler) GetAllAttendances(c *fiber.Ctx) error {
	attendances, err := h.service.GetAllAttendances(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(attendances)
}

func (h *handler) GetAttendanceByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid attendance ID"})
	}
	attendance, err := h.service.GetAttendanceByID(context.Background(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(attendance)
}

func (h *handler) UpdateAttendance(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid attendance ID"})
	}

	req := UpdateAttendanceRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.ID = id
	req.UpdatedBy = "system" // replace with actual user
	// req.UpdatedAt = time.Now()

	if err := h.service.UpdateAttendance(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Attendance updated successfully"})
}

func (h *handler) DeleteAttendance(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid attendance ID"})
	}

	if err := h.service.DeleteAttendance(context.Background(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Attendance deleted successfully"})
}
