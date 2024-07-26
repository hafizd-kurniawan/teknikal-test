package employee

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service
}

func newHandler(service service) handler {
	return handler{
		service: service,
	}
}

func (h *handler) CreateEmployee(c *fiber.Ctx) error {
	req := CreateEmployeeRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.CreatedBy = "system" // replace with actual user
	// req.CreatedAt = time.Now()

	if err := h.service.CreateEmployee(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Employee created successfully"})
}

func (h *handler) GetAllEmployees(c *fiber.Ctx) error {
	employees, err := h.service.GetAllEmployees(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(employees)
}

func (h *handler) GetEmployeeByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
	}
	employee, err := h.service.GetEmployeeByID(context.Background(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(employee)
}

func (h *handler) UpdateEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
	}

	req := UpdateEmployeeRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	req.EmployeeID = id
	req.UpdatedBy = "system" // replace with actual user
	// req.UpdatedAt = time.Now()

	if err := h.service.UpdateEmployee(context.Background(), req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Employee updated successfully"})
}

func (h *handler) DeleteEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
	}

	if err := h.service.DeleteEmployee(context.Background(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Employee deleted successfully"})
}
