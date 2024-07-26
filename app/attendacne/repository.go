package attendacne

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{db: db}
}

func (r *repository) CreateAttendance(ctx context.Context, att Attendance) error {
	query := `INSERT INTO attendance (employee_id, location_id, absent_in, absent_out, created_at, created_by, updated_at, updated_by)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.ExecContext(ctx, query, att.EmployeeID, att.LocationID, att.AbsentIn, att.AbsentOut, att.CreatedAt, att.CreatedBy, att.UpdatedAt, att.UpdatedBy)
	return err
}

func (r *repository) UpdateAttendance(ctx context.Context, att Attendance) error {
	query := `UPDATE attendance SET employee_id = $1, location_id = $2, absent_in = $3, absent_out = $4, updated_at = $5, updated_by = $6
	          WHERE attendance_id = $7`
	_, err := r.db.ExecContext(ctx, query, att.EmployeeID, att.LocationID, att.AbsentIn, att.AbsentOut, att.UpdatedAt, att.UpdatedBy, att.AttendanceID)
	return err
}

func (r *repository) UpdateCheckoutAttendance(ctx context.Context, att Attendance) error {
	var existingAbsentOut sql.NullTime // untuk menyimpan nilai absen
	checkQuery := `SELECT absent_out FROM employees WHERE id = $1`
	err := r.db.GetContext(ctx, &existingAbsentOut, checkQuery, att.AttendanceID)
	if err != nil {
		return err
	}

	if existingAbsentOut.Valid {
		return errors.New("absent_out is already set for this employee")
	}

	updateQuery := `UPDATE employees 
                    SET absent_out = :absent_out, updated_at = :updated_at, updated_by = :updated_by 
                    WHERE id = :id`

	_, err = r.db.NamedExecContext(ctx, updateQuery, map[string]interface{}{
		"id":         att.AttendanceID,
		"absent_out": time.Now(),
		"updated_at": time.Now(),
		"updated_by": att.UpdatedBy,
	})
	return err
}

func (r *repository) DeleteAttendance(ctx context.Context, id int) error {
	query := `UPDATE attendance SET deleted_at = $1 WHERE attendance_id = $2`
	_, err := r.db.ExecContext(ctx, query, time.Now(), id)
	return err
}

func (r *repository) GetAttendanceByID(ctx context.Context, id int) (Attendance, error) {
	var att Attendance
	query := `SELECT * FROM attendance WHERE attendance_id = $1 AND deleted_at IS NULL`
	err := r.db.GetContext(ctx, &att, query, id)
	if err != nil {
		return Attendance{}, err
	}
	return att, nil
}

func (r *repository) GetAllAttendances(ctx context.Context) ([]Attendance, error) {
	var attendances []Attendance
	query := `SELECT * FROM attendance WHERE deleted_at IS NULL`
	err := r.db.SelectContext(ctx, &attendances, query)
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (r *repository) GetAttendanceReport(ctx context.Context, startDate, endDate time.Time) ([]AttendanceReport, error) {
	var report []AttendanceReport
	query := `SELECT a.attendance_id, e.employee_code, e.employee_name, d.department_name, p.position_name, l.location_name, 
	                 a.absent_in, a.absent_out
	          FROM attendance a
	          JOIN employees e ON a.employee_id = e.employee_id
	          JOIN departments d ON e.department_id = d.department_id
	          JOIN positions p ON e.position_id = p.position_id
	          JOIN locations l ON a.location_id = l.location_id
	          WHERE a.absent_in >= $1 AND a.absent_out <= $2 AND a.deleted_at IS NULL`
	err := r.db.SelectContext(ctx, &report, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return report, nil
}
