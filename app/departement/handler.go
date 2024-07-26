package departement

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(service service) handler {
	return handler{svc: service}
}
func (h *handler) CreateDepartment(c *fiber.Ctx) error {
	req := CreateDepartmentRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.CreatedBy = "system" // replace with actual user

	if err := h.svc.CreateDepartment(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Department created successfully"})
}

func (h *handler) GetAllDepartments(c *fiber.Ctx) error {
	departments, err := h.svc.GetAllDepartments(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(departments)
}

func (h *handler) GetDepartmentByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}
	department, err := h.svc.GetDepartmentByID(context.Background(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(department)
}

func (h *handler) UpdateDepartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}

	req := UpdateDepartmentRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.DepartmentID = id
	req.UpdatedBy = "system" // replace with actual user

	if err := h.svc.UpdateDepartment(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Department updated successfully"})
}

func (h *handler) DeleteDepartment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid department ID"})
	}

	if err := h.svc.DeleteDepartment(context.Background(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Department deleted successfully"})
}
