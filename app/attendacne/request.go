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
	EmployeeID int    `json:"employee_id" validate:"required"`
	LocationID int    `json:"location_id" validate:"required"`
	CreatedBy  string `json:"created_by" validate:"required"`
}
type UpdateAttendanceRequest struct {
	AttendanceID int       `json:"id" validate:"required"`
	EmployeeID   int       `json:"employee_id" validate:"required"`
	LocationID   int       `json:"location_id" validate:"required"`
	AbsentIn     time.Time `json:"absent_in" validate:"required"`
	AbsentOut    time.Time `json:"absent_out" validate:"required"`
	UpdatedBy    string    `json:"updated_by" validate:"required"`
}
