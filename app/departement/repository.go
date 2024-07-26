package departement

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}
func (r *repository) CreateDepartment(ctx context.Context, dept Department) error {
	query := `INSERT INTO departments (department_name, created_at, created_by, updated_at, updated_by) 
	          VALUES (:department_name, :created_at, :created_by, :updated_at, :updated_by)`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"department_name": dept.DepartmentName,
		"created_at":      time.Now(),
		"created_by":      dept.CreatedBy,
		"updated_at":      time.Now(),
		"updated_by":      dept.UpdatedBy,
	})
	return err
}

func (r *repository) GetDepartmentByID(ctx context.Context, id int) (Department, error) {
	query := `SELECT department_id, department_name, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM departments WHERE department_id = $1 AND deleted_at IS NULL`
	var dept Department
	err := r.db.GetContext(ctx, &dept, query, id)
	return dept, err
}

func (r *repository) GetAllDepartments(ctx context.Context) ([]Department, error) {
	query := `SELECT department_id, department_name, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM departments WHERE deleted_at IS NULL`
	var departments []Department
	err := r.db.SelectContext(ctx, &departments, query)
	return departments, err
}

func (r *repository) UpdateDepartment(ctx context.Context, dept Department) error {
	query := `UPDATE departments SET department_name = :department_name, updated_at = :updated_at, updated_by = :updated_by 
	          WHERE department_id = :department_id AND deleted_at IS NULL`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"department_name": dept.DepartmentName,
		"updated_at":      time.Now(),
		"updated_by":      dept.UpdatedBy,
		"department_id":   dept.DepartmentID,
	})
	return err
}

func (r *repository) DeleteDepartment(ctx context.Context, id int) error {
	query := `UPDATE departments SET deleted_at = :deleted_at WHERE department_id = :department_id`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"deleted_at":    time.Now(),
		"department_id": id,
	})
	return err
}
