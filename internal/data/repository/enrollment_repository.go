package repository

import (
	"context"
	"database/sql"
	"ppugenrollment/internal/data/sqlcgen"
	"ppugenrollment/pkg/domain"
)

type DefaultEnrollmentRepository struct {
	db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) *DefaultEnrollmentRepository {
	return &DefaultEnrollmentRepository{db}
}

func (d *DefaultEnrollmentRepository) InsertEnrollment(application *domain.EnrollmentApplication) (int,
	*domain.AppError) {
	queries := sqlcgen.New(d.db)

	result, err := queries.CreateEnrollmentApplication(context.Background(), sqlcgen.CreateEnrollmentApplicationParams{
		Student:  int32(application.Student.ID),
		Project:  int32(application.Project.ID),
		Schedule: int32(application.Schedule),
	})

	if err != nil {
		return 0, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	lastInsertedID, _ := result.LastInsertId()

	return int(lastInsertedID), nil
}
