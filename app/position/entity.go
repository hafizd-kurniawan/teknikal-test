package position

import (
	"errors"
	"time"
)

type Position struct {
	PositionID   int        `json:"id" db:"position_id"`
	DepartmentID int        `json:"department_id" db:"department_id"`
	PositionName string     `json:"name" db:"position_name"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	CreatedBy    string     `json:"created_by" db:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy    string     `json:"updated_by" db:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func NewPosition(req CreatePositionRequest) Position {
	pos := Position{
		DepartmentID: req.DepartmentID,
		PositionName: req.PositionName,
		CreatedAt:    time.Now(),
		CreatedBy:    req.CreatedBy,
		UpdatedAt:    time.Now(),
		UpdatedBy:    req.CreatedBy,
	}
	return pos
}

func NewUpdatePosition(req UpdatePositionRequest) Position {
	pos := Position{
		PositionID:   req.PositionID,
		DepartmentID: req.DepartmentID,
		PositionName: req.PositionName,
		UpdatedBy:    req.UpdatedBy,
		UpdatedAt:    time.Now(),
	}
	return pos
}

func (p *Position) Validate() error {
	if p.PositionName == "" {
		return errors.New("position name is required")
	}
	if p.DepartmentID == 0 {
		return errors.New("department id is required")
	}
	return nil
}
