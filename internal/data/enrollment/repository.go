package enrollment

import (
	"github.com/jmoiron/sqlx"
	"ppugenrollment/internal/domain"
)

type DefaultRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DefaultRepository {
	return &DefaultRepository{db}
}

func (d *DefaultRepository) InsertEnrollment(application *domain.EnrollmentApplication) (int, *domain.AppError) {
	applicationModel := fromEnrollmentApplicationToModel(application)

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
