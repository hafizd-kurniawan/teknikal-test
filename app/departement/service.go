package departement

import (
	"context"
)

type IDepartmentRepository interface {
	CreateDepartment(ctx context.Context, department Department) error
	GetDepartmentByID(ctx context.Context, id int) (Department, error)
	GetAllDepartments(ctx context.Context) ([]Department, error)
	UpdateDepartment(ctx context.Context, department Department) error
	DeleteDepartment(ctx context.Context, id int) error
}

type service struct {
	repo IDepartmentRepository
}

func newService(repo IDepartmentRepository) service {
	return service{repo: repo}
}

func (s *service) CreateDepartment(ctx context.Context, req CreateDepartmentRequest) error {
	departement := NewDepartment(req)

	if err := departement.Validate(); err != nil {
		return err
	}

	return s.repo.CreateDepartment(ctx, departement)
}

func (s *service) GetDepartmentByID(ctx context.Context, id int) (Department, error) {
	return s.repo.GetDepartmentByID(ctx, id)
}

func (s *service) GetAllDepartments(ctx context.Context) ([]Department, error) {
	return s.repo.GetAllDepartments(ctx)
}

func (s *service) UpdateDepartment(ctx context.Context, req UpdateDepartmentRequest) error {
	dept := NewUpdateDepartment(req)
	return s.repo.UpdateDepartment(ctx, dept)
}

func (s *service) DeleteDepartment(ctx context.Context, id int) error {
	return s.repo.DeleteDepartment(ctx, id)
}
