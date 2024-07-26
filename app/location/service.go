package location

import (
	"context"
)

type ILocationRepository interface {
	CreateLocation(ctx context.Context, location Location) error
	GetLocationByID(ctx context.Context, id int) (Location, error)
	GetAllLocations(ctx context.Context) ([]Location, error)
	UpdateLocation(ctx context.Context, location Location) error
	DeleteLocation(ctx context.Context, id int) error
}

type service struct {
	repo ILocationRepository
}

func newService(repo ILocationRepository) service {
	return service{repo: repo}
}
func (s *service) CreateLocation(ctx context.Context, req CreateLocationRequest) error {
	loc := NewLocation(req)
	if err := loc.Validate(); err != nil {
		return err
	}
	return s.repo.CreateLocation(ctx, loc)
}

func (s *service) GetLocationByID(ctx context.Context, id int) (Location, error) {
	return s.repo.GetLocationByID(ctx, id)
}

func (s *service) GetAllLocations(ctx context.Context) ([]Location, error) {
	return s.repo.GetAllLocations(ctx)
}

func (s *service) UpdateLocation(ctx context.Context, req UpdateLocationRequest) error {
	loc := NewUpdateLocation(req)
	return s.repo.UpdateLocation(ctx, loc)
}

func (s *service) DeleteLocation(ctx context.Context, id int) error {
	return s.repo.DeleteLocation(ctx, id)
}
