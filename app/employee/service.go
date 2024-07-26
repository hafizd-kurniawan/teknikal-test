package employee

import (
	"context"
)

type IRepository interface {
	CreateEmployee(ctx context.Context, employee Employee) error
	GetEmployeeByID(ctx context.Context, id int) (Employee, error)
	GetAllEmployees(ctx context.Context) ([]Employee, error)
	UpdateEmployee(ctx context.Context, employee Employee) error
	DeleteEmployee(ctx context.Context, id int) error
	// Login(ctx context.Context, email, password string) (Employee, error)
	// ForgotPassword(ctx context.Context, email string) error
	// GetEmployeeByNameAndDeptIdAndPositionId()
	// GetEmployeeByCode()
	// GetEmployeeByName()
	// GetEmployeeById()
	// GetAllEmployee()
	// ForgotPass()
	// Register()
	// Delete()
	// Insert(ctx context.Context, model EmployeeEntity) (err error)
	// Update()
	// Login(ctx context.Context, model EmployeeEntity) (err error)
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
	return s.repo.CreateEmployee(ctx, emp)
}

func (s *service) UpdateEmployee(ctx context.Context, req UpdateEmployeeRequest) error {
	emp := NewUpdateRequest(req)
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
