package location

import (
	"errors"
	"time"
)

type Location struct {
	LocationId   int        `json:"id" db:"location_id"`
	LocationName string     `json:"name" db:"location_name"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	CreatedBy    string     `json:"created_by" db:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy    string     `json:"updated_by" db:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func NewLocation(req CreateLocationRequest) Location {
	loc := Location{
		LocationName: req.LocationName,
		CreatedAt:    time.Now(),
		CreatedBy:    req.CreatedBy,
		UpdatedAt:    time.Now(),
		UpdatedBy:    req.CreatedBy,
	}
	return loc
}

func NewUpdateLocation(req UpdateLocationRequest) Location {
	loc := Location{
		LocationId:   req.LocationId,
		LocationName: req.LocationName,
		UpdatedBy:    req.UpdatedBy,
	}
	return loc
}

func (l *Location) Validate() error {
	if l.LocationName == "" {
		return errors.New("location name is required")
	}
	return nil
}
