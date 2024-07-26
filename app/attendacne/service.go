package attendacne

import (
	"context"
	"time"
)

type IAttendanceRepository interface {
	CreateAttendance(ctx context.Context, attendance Attendance) error
	GetAttendanceByID(ctx context.Context, id int) (Attendance, error)
	GetAllAttendances(ctx context.Context) ([]Attendance, error)
	UpdateAttendance(ctx context.Context, attendance Attendance) error
	DeleteAttendance(ctx context.Context, id int) error
	GetAttendanceReport(ctx context.Context, startTime, endTime time.Time) ([]AttendanceReport, error)
}

type service struct {
	repo IAttendanceRepository
}

func newService(repo IAttendanceRepository) service {
	return service{repo: repo}
}
func (s *service) CreateAttendance(ctx context.Context, req CreateAttendanceRequest) error {
	att := Attendance{
		EmployeeID: req.EmployeeID,
		LocationID: req.LocationID,
		AbsentIn:   req.AbsentIn,
		AbsentOut:  req.AbsentOut,
		CreatedAt:  time.Now(),
		CreatedBy:  req.CreatedBy,
		UpdatedAt:  time.Now(),
		UpdatedBy:  req.CreatedBy,
	}

	return s.repo.CreateAttendance(ctx, att)
}

func (s *service) GetAttendanceByID(ctx context.Context, id int) (Attendance, error) {
	return s.repo.GetAttendanceByID(ctx, id)
}

func (s *service) GetAllAttendances(ctx context.Context) ([]Attendance, error) {
	return s.repo.GetAllAttendances(ctx)
}

func (s *service) UpdateAttendance(ctx context.Context, req UpdateAttendanceRequest) error {
	att := Attendance{
		ID:         req.ID,
		EmployeeID: req.EmployeeID,
		LocationID: req.LocationID,
		AbsentIn:   req.AbsentIn,
		AbsentOut:  req.AbsentOut,
		UpdatedBy:  req.UpdatedBy,
		UpdatedAt:  time.Now(),
	}

	return s.repo.UpdateAttendance(ctx, att)
}

func (s *service) DeleteAttendance(ctx context.Context, id int) error {
	return s.repo.DeleteAttendance(ctx, id)
}

func (s *service) GetAttendanceReport(ctx context.Context, startDate, endDate time.Time) ([]AttendanceReport, error) {
	return s.repo.GetAttendanceReport(ctx, startDate, endDate)
}
