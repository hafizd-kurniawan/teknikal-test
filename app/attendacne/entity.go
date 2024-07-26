package attendacne

import (
	"errors"
	"time"
)

type Attendance struct {
	AttendanceID int        `json:"id" db:"attendance_id"`
	EmployeeID   int        `json:"employee_id" db:"employee_id"`
	LocationID   int        `json:"location_id" db:"location_id"`
	AbsentIn     time.Time  `json:"absent_in" db:"absent_in"`
	AbsentOut    time.Time  `json:"absent_out" db:"absent_out"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	CreatedBy    string     `json:"created_by" db:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy    string     `json:"updated_by" db:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (a *Attendance) Validate() error {
	if a.EmployeeID == 0 {
		return errors.New("employee id is required")
	}
	if a.LocationID == 0 {
		return errors.New("location id is required")
	}
	// if a.AbsentIn.IsZero() {
	// 	return errors.New("absent in time is required")
	// }
	// if a.AbsentOut.IsZero() {
	// 	return errors.New("absent out time is required")
	// }
	return nil
}

func NewAttendance(req CreateAttendanceRequest) Attendance {
	att := Attendance{
		EmployeeID: req.EmployeeID,
		LocationID: req.LocationID,
		AbsentIn:   time.Now(),
		AbsentOut:  time.Now(),
		CreatedAt:  time.Now(),
		CreatedBy:  req.CreatedBy,
		UpdatedAt:  time.Now(),
		UpdatedBy:  req.CreatedBy,
	}
	return att
}

func NewUpdateAttendance(req UpdateAttendanceRequest) Attendance {
	att := Attendance{
		AttendanceID: req.AttendanceID,
		EmployeeID:   req.EmployeeID,
		LocationID:   req.LocationID,
		AbsentIn:     req.AbsentIn,
		AbsentOut:    req.AbsentOut,
		UpdatedBy:    req.UpdatedBy,
		UpdatedAt:    time.Now(),
	}
	return att
}
