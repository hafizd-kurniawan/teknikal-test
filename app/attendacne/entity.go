package attendacne

import (
	"errors"
	"time"
)

type Attendance struct {
	ID         int        `json:"id" db:"attendance_id"`
	EmployeeID int        `json:"employee_id" db:"employee_id"`
	LocationID int        `json:"location_id" db:"location_id"`
	AbsentIn   time.Time  `json:"absent_in" db:"absent_in"`
	AbsentOut  time.Time  `json:"absent_out" db:"absent_out"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	CreatedBy  string     `json:"created_by" db:"created_by"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy  string     `json:"updated_by" db:"updated_by"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (a *Attendance) Validate() error {
	if a.EmployeeID == 0 {
		return errors.New("employee id is required")
	}
	if a.LocationID == 0 {
		return errors.New("location id is required")
	}
	if a.AbsentIn.IsZero() {
		return errors.New("absent in time is required")
	}
	if a.AbsentOut.IsZero() {
		return errors.New("absent out time is required")
	}
	return nil
}

func NewAttendance(employeeID, locationID int, absentIn, absentOut time.Time, createdBy string) (*Attendance, error) {
	att := &Attendance{
		EmployeeID: employeeID,
		LocationID: locationID,
		AbsentIn:   absentIn,
		AbsentOut:  absentOut,
		CreatedAt:  time.Now(),
		CreatedBy:  createdBy,
		UpdatedAt:  time.Now(),
		UpdatedBy:  createdBy,
	}
	if err := att.Validate(); err != nil {
		return nil, err
	}
	return att, nil
}
