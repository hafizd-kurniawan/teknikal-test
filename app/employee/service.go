package employee

import (
	"context"
	"errors"
)

type IRepository interface {
	CreateEmployee(ctx context.Context, employee Employee) error
	GetEmployeeByID(ctx context.Context, id int) (Employee, error)
	GetAllEmployees(ctx context.Context) ([]Employee, error)
	UpdateEmployee(ctx context.Context, employee Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	GetEmployeeByCode(ctx context.Context, code string) (Employee, error)
}

type service struct {
	repo IRepository
}

func newService(repository IRepository) service {
	return service{repo: repository}
}

func (s *service) CreateEmployee(ctx context.Context, req CreateEmployeeRequest) error {
	emp := NewCreateEmployee(req)
	if err := emp.Validate(); err != nil {
		return err
	}

	if err := emp.EncryptPassword(); err != nil {
		return err
	}
	if _, err := s.repo.GetEmployeeByCode(ctx, req.EmployeeCode); err == nil {
		return errors.New("duplicate data")
	}

	return s.repo.CreateEmployee(ctx, emp)
}

func (s *service) UpdateEmployee(ctx context.Context, req UpdateEmployeeRequest) error {
	emp := NewUpdateRequest(req)

	if err := emp.EncryptPassword(); err != nil {
		return err
	}

	return s.repo.UpdateEmployee(ctx, emp)
}

func (s *service) GetEmployeeByID(ctx context.Context, id int) (Employee, error) {
	return s.repo.GetEmployeeByID(ctx, id)
}

func (s *service) GetAllEmployees(ctx context.Context) ([]Employee, error) {
	return s.repo.GetAllEmployees(ctx)
}

func (s *service) DeleteEmployee(ctx context.Context, id int) error {
	return s.repo.DeleteEmployee(ctx, id)
}

func (s *service) Login(ctx context.Context, code, password string) (string, error) {
	var token string
	employee, err := s.repo.GetEmployeeByCode(ctx, code)
	if err != nil {
		return "", err
	}

	if err := employee.ValidatePasswordFromPlainText(password, employee.Password); err != nil {
		return "", err
	}

	token, err = employee.GenerateAccessToken()
	if err != nil {
		return "", err
	}

	return token, nil
}
