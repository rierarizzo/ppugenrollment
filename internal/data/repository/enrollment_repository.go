package repository

import (
	"github.com/jmoiron/sqlx"
	"ppugenrollment/internal/data/mappers"
	"ppugenrollment/internal/domain"
)

type DefaultEnrollmentRepository struct {
	db *sqlx.DB
}

func NewEnrollmentRepository(db *sqlx.DB) *DefaultEnrollmentRepository {
	return &DefaultEnrollmentRepository{db}
}

// InsertEnrollment takes an EnrollmentApplication and inserts it into the enrollment_application table in the database.
// It returns the ID of the last inserted row and an AppError if any error occurred during the insertion.
func (d *DefaultEnrollmentRepository) InsertEnrollment(application *domain.EnrollmentApplication) (
	int,
	*domain.AppError) {
	applicationModel := mappers.FromEnrollmentApplicationToModel(application)

	insertApplicationInSchema := `
		INSERT INTO enrollment_application (student, project, schedule) VALUES (?,?,?)
	`
	result, err := d.db.Exec(
		insertApplicationInSchema,
		applicationModel.Student,
		applicationModel.Project,
		applicationModel.Schedule)
	if err != nil {
		return 0, domain.NewAppError(err, domain.RepositoryError)
	}

	lastInsertedID, _ := result.LastInsertId()

	return int(lastInsertedID), nil
}
