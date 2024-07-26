package departement

import (
	"errors"
	"time"
)

type Department struct {
	DepartmentID   int        `json:"id" db:"department_id"`
	DepartmentName string     `json:"name" db:"department_name"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	CreatedBy      string     `json:"created_by" db:"created_by"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy      string     `json:"updated_by" db:"updated_by"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func NewDepartment(req CreateDepartmentRequest) Department {
	dept := Department{
		DepartmentName: req.DepartmentName,
		CreatedAt:      time.Now(),
		CreatedBy:      req.CreatedBy,
		UpdatedAt:      time.Now(),
		UpdatedBy:      req.CreatedBy,
	}
	return dept
}

func NewUpdateDepartment(req UpdateDepartmentRequest) Department {
	dept := Department{
		DepartmentID:   req.DepartmentID,
		DepartmentName: req.DepartmentName,
		UpdatedBy:      req.UpdatedBy,
		UpdatedAt:      time.Now(),
	}
	return dept
}

func (d *Department) Validate() error {
	if d.DepartmentName == "" {
		return errors.New("department name is required")
	}
	return nil
}
