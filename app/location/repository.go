package location

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
func (r *repository) CreateLocation(ctx context.Context, loc Location) error {
	query := `INSERT INTO locations (location_name, created_at, created_by, updated_at, updated_by) 
	          VALUES (:location_name, :created_at, :created_by, :updated_at, :updated_by)`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"location_name": loc.LocationName,
		"created_at":    time.Now(),
		"created_by":    loc.CreatedBy,
		"updated_at":    time.Now(),
		"updated_by":    loc.UpdatedBy,
	})
	return err
}

func (r *repository) GetLocationByID(ctx context.Context, id int) (Location, error) {
	query := `SELECT location_id, location_name, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM locations WHERE location_id = $1 AND deleted_at IS NULL`
	var loc Location
	err := r.db.GetContext(ctx, &loc, query, id)
	return loc, err
}

func (r *repository) GetAllLocations(ctx context.Context) ([]Location, error) {
	query := `SELECT location_id, location_name, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM locations WHERE deleted_at IS NULL`
	var locations []Location
	err := r.db.SelectContext(ctx, &locations, query)
	return locations, err
}

func (r *repository) UpdateLocation(ctx context.Context, loc Location) error {
	query := `UPDATE locations SET location_name = :location_name, updated_at = :updated_at, updated_by = :updated_by 
	          WHERE location_id = :location_id AND deleted_at IS NULL`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"location_name": loc.LocationName,
		"updated_at":    time.Now(),
		"updated_by":    loc.UpdatedBy,
		"location_id":   loc.LocationId,
	})
	return err
}

func (r *repository) DeleteLocation(ctx context.Context, id int) error {
	query := `UPDATE locations SET deleted_at = :deleted_at WHERE location_id = :location_id`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"deleted_at":  time.Now(),
		"location_id": id,
	})
	return err
}
