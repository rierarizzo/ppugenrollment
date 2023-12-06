package approval

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
	"ppugenrollment/internal/domain"
)

type DefaultRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DefaultRepository {
	return &DefaultRepository{db}
}

func (d *DefaultRepository) ApproveEnrollmentApplication(applicationID, approvedBy int) (int, *domain.AppError) {
	tx, _ := d.db.Beginx()

	updateEnrollmentApplication := `
		UPDATE enrollment_application SET status='A' WHERE id=?
	`
	result, err := tx.Exec(updateEnrollmentApplication, applicationID)
	if err != nil {
		_ = tx.Rollback()
		return 0, domain.NewAppError(err, domain.RepositoryError)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected != 1 {
		slog.Error("0 rows affected")
		_ = tx.Rollback()
		return 0, domain.NewAppErrorWithType(domain.RepositoryError)
	}

	insertIntoEnrollmentGenerated := `
		INSERT INTO enrollment_generated (enrollment_application, approved_by) VALUES (?,?)
	`
	result, err = tx.Exec(insertIntoEnrollmentGenerated, applicationID, approvedBy)
	if err != nil {
		_ = tx.Rollback()
		return 0, domain.NewAppError(err, domain.RepositoryError)
	}

	generatedID, _ := result.LastInsertId()

	return int(generatedID), nil
}

func (d *DefaultRepository) SelectEnrollmentGenerated(generatedID int) (*domain.EnrollmentGenerated, *domain.AppError) {
	var generatedModel EnrollmentGeneratedModel

	enrollmentGeneratedQuery := `
		SELECT 
			eg.id as id,
			eg.enrollment_application as application_id,
			eg.generated_at as generated_at,
			p.id as project_id,
			p.description as project_description,
			p.schedule as project_schedule,
			p.starts as project_starts,
			p.ends as project_ends,
			c.id as company_id,
			c.name as company_name,
			c.ruc as company_ruc,
			su.id as approver_id,
			su.id_card_number as approver_card_number,
			su.name as approver_name,
			su.surname as approver_surname
		FROM enrollment_generated eg 
		    INNER JOIN enrollment_application ea ON eg.enrollment_application = ea.id
			INNER JOIN project p ON ea.project = p.id 
		    INNER JOIN company c ON p.company = c.id 
		    INNER JOIN sys_user su ON eg.approved_by = su.id 
		WHERE eg.id=?
	`
	err := d.db.Get(&generatedModel, enrollmentGeneratedQuery, generatedID)
	if err != nil {
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	generated := fromModelToEnrollmentGenerated(&generatedModel)

	return &generated, nil
}
