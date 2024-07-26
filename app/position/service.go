package position

import (
	"context"
)

type IPositionRepository interface {
	CreatePosition(ctx context.Context, position Position) error
	GetPositionByID(ctx context.Context, id int) (Position, error)
	GetAllPositions(ctx context.Context) ([]Position, error)
	UpdatePosition(ctx context.Context, position Position) error
	DeletePosition(ctx context.Context, id int) error
}
type service struct {
	repo IPositionRepository
}

func newService(repo IPositionRepository) service {
	return service{repo: repo}
}

func (s *service) CreatePosition(ctx context.Context, req CreatePositionRequest) error {
	pos := NewPosition(req)
	if err := pos.Validate(); err != nil {
		return err
	}
	return s.repo.CreatePosition(ctx, pos)
}

func (s *service) GetPositionByID(ctx context.Context, id int) (Position, error) {
	return s.repo.GetPositionByID(ctx, id)
}

func (s *service) GetAllPositions(ctx context.Context) ([]Position, error) {
	return s.repo.GetAllPositions(ctx)
}

func (s *service) UpdatePosition(ctx context.Context, req UpdatePositionRequest) error {
	pos := NewUpdatePosition(req)
	return s.repo.UpdatePosition(ctx, pos)
}

func (s *service) DeletePosition(ctx context.Context, id int) error {
	return s.repo.DeletePosition(ctx, id)
}
