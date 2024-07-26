package employee

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) (repo repository) {
	return repository{db: db}
}

func (r *repository) GetEmployeeByCode(ctx context.Context, code string) (Employee, error) {
	var emp Employee
	query := `SELECT * FROM employees WHERE employee_code = $1 AND deleted_at IS NULL`
	err := r.db.GetContext(ctx, &emp, query, code)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

func (r *repository) CreateEmployee(ctx context.Context, emp Employee) error {
	query := `INSERT INTO employees (employee_code, employee_name, password, department_id, position_id, superior, created_at, created_by, updated_at, updated_by) 
	          VALUES (:employee_code, :employee_name, :password, :department_id, :position_id, :superior, :created_at, :created_by, :updated_at, :updated_by)`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"employee_code": emp.EmployeeCode,
		"employee_name": emp.EmployeeName,
		"password":      emp.Password,
		"department_id": emp.DepartmentID,
		"position_id":   emp.PositionID,
		"superior":      emp.Superior,
		"created_at":    time.Now(),
		"created_by":    emp.CreatedBy,
		"updated_at":    time.Now(),
		"updated_by":    emp.UpdatedBy,
	})
	return err
}

func (r *repository) GetEmployeeByID(ctx context.Context, id int) (Employee, error) {
	query := `SELECT employee_id, employee_code, employee_name, password, department_id, position_id, superior, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM employees WHERE employee_id = $1 AND deleted_at IS NULL`
	var emp Employee
	err := r.db.GetContext(ctx, &emp, query, id)
	return emp, err
}

func (r *repository) GetAllEmployees(ctx context.Context) ([]Employee, error) {
	query := `SELECT employee_id, employee_code, employee_name, password, department_id, position_id, superior, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM employees WHERE deleted_at IS NULL`
	var employees []Employee
	err := r.db.SelectContext(ctx, &employees, query)
	return employees, err
}

func (r *repository) UpdateEmployee(ctx context.Context, emp Employee) error {
	query := `UPDATE employees SET employee_code = :employee_code, employee_name = :employee_name, password = :password, 
	          department_id = :department_id, position_id = :position_id, superior = :superior, updated_at = :updated_at, 
	          updated_by = :updated_by WHERE employee_id = :employee_id AND deleted_at IS NULL`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"employee_code": emp.EmployeeCode,
		"employee_name": emp.EmployeeName,
		"password":      emp.Password,
		"department_id": emp.DepartmentID,
		"position_id":   emp.PositionID,
		"superior":      emp.Superior,
		"updated_at":    time.Now(),
		"updated_by":    emp.UpdatedBy,
		"employee_id":   emp.EmployeeId,
	})
	return err
}

func (r *repository) DeleteEmployee(ctx context.Context, id int) error {
	query := `UPDATE employees SET deleted_at = :deleted_at WHERE employee_id = :employee_id`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"deleted_at":  time.Now(),
		"employee_id": id,
	})
	return err
}
