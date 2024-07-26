package position

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
func (r *repository) CreatePosition(ctx context.Context, pos Position) error {
	query := `INSERT INTO positions (department_id, position_name, created_at, created_by, updated_at, updated_by) 
	          VALUES (:department_id, :position_name, :created_at, :created_by, :updated_at, :updated_by)`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"department_id": pos.DepartmentID,
		"position_name": pos.PositionName,
		"created_at":    time.Now(),
		"created_by":    pos.CreatedBy,
		"updated_at":    time.Now(),
		"updated_by":    pos.UpdatedBy,
	})
	return err
}

func (r *repository) GetPositionByID(ctx context.Context, id int) (Position, error) {
	query := `SELECT position_id, department_id, position_name, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM positions WHERE position_id = $1 AND deleted_at IS NULL`
	var pos Position
	err := r.db.GetContext(ctx, &pos, query, id)
	return pos, err
}

func (r *repository) GetAllPositions(ctx context.Context) ([]Position, error) {
	query := `SELECT position_id, department_id, position_name, created_at, created_by, updated_at, updated_by, deleted_at 
	          FROM positions WHERE deleted_at IS NULL`
	var positions []Position
	err := r.db.SelectContext(ctx, &positions, query)
	return positions, err
}

func (r *repository) UpdatePosition(ctx context.Context, pos Position) error {
	query := `UPDATE positions SET department_id = :department_id, position_name = :position_name, updated_at = :updated_at, updated_by = :updated_by 
	          WHERE position_id = :position_id AND deleted_at IS NULL`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"department_id": pos.DepartmentID,
		"position_name": pos.PositionName,
		"updated_at":    time.Now(),
		"updated_by":    pos.UpdatedBy,
		"position_id":   pos.PositionID,
	})
	return err
}

func (r *repository) DeletePosition(ctx context.Context, id int) error {
	query := `UPDATE positions SET deleted_at = :deleted_at WHERE position_id = :position_id`
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"deleted_at":  time.Now(),
		"position_id": id,
	})
	return err
}
