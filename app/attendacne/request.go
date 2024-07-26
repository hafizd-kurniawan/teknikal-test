package attendacne

import "time"

type AttendanceReport struct {
	Date           time.Time `json:"date"`
	EmployeeCode   string    `json:"employee_code"`
	EmployeeName   string    `json:"employee_name"`
	DepartmentName string    `json:"department_name"`
	PositionName   string    `json:"position_name"`
	LocationName   string    `json:"location_name"`
	AbsentIn       time.Time `json:"absent_in"`
	AbsentOut      time.Time `json:"absent_out"`
}
type CreateAttendanceRequest struct {
	EmployeeID int       `json:"employee_id" validate:"required"` // Employee ID
	LocationID int       `json:"location_id" validate:"required"` // Location ID
	AbsentIn   time.Time `json:"absent_in" validate:"required"`   // Time of absence start
	AbsentOut  time.Time `json:"absent_out" validate:"required"`  // Time of absence end
	CreatedBy  string    `json:"created_by" validate:"required"`  // User who created the record
}
type UpdateAttendanceRequest struct {
	ID         int       `json:"id" validate:"required"`          // Attendance ID
	EmployeeID int       `json:"employee_id" validate:"required"` // Employee ID
	LocationID int       `json:"location_id" validate:"required"` // Location ID
	AbsentIn   time.Time `json:"absent_in" validate:"required"`   // Time of absence start
	AbsentOut  time.Time `json:"absent_out" validate:"required"`  // Time of absence end
	UpdatedBy  string    `json:"updated_by" validate:"required"`  // User who updated the record
}
