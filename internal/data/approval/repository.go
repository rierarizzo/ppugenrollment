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

func (d *DefaultRepository) ApproveEnrollmentApplication(applicationID, approvedBy int) (
	*domain.EnrollmentGenerated,
	*domain.AppError) {
	tx, _ := d.db.Beginx()

	updateEnrollmentApplication := `
		UPDATE enrollment_application SET status='A' WHERE id=?
	`
	result, err := tx.Exec(updateEnrollmentApplication, applicationID)
	if err != nil {
		_ = tx.Rollback()
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected != 1 {
		slog.Error("0 rows affected")
		_ = tx.Rollback()
		return nil, domain.NewAppError("application already approved", domain.RepositoryError)
	}

	insertIntoEnrollmentGenerated := `
		INSERT INTO enrollment_generated (enrollment_application, approved_by) VALUES (?,?)
	`
	result, err = tx.Exec(insertIntoEnrollmentGenerated, applicationID, approvedBy)
	if err != nil {
		_ = tx.Rollback()
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	generatedID, _ := result.LastInsertId()

	var generatedModel EnrollmentGeneratedModel

	enrollmentGeneratedQuery := `
		SELECT eg.id AS id,
			eg.enrollment_application AS application_id,
			eg.generated_at AS generated_at,
			p.id AS project_id,
			p.description AS project_description,
			ps.schedule AS project_schedule,
			p.starts AS project_starts,
			p.ends AS project_ends,
			c.id AS company_id,
			c.name AS company_name,
			c.ruc AS company_ruc,
			su.id AS approver_id,
			su.id_card_number AS approver_card_number,
			su.name AS approver_name,
			su.surname AS approver_surname
		FROM enrollment_generated eg 
		    INNER JOIN enrollment_application ea ON eg.enrollment_application = ea.id
			INNER JOIN project p ON ea.project = p.id 
		    INNER JOIN project_schedule ps ON ea.schedule = ps.id 
		    INNER JOIN company c ON p.company = c.id 
		    INNER JOIN user su ON eg.approved_by = su.id 
		WHERE eg.id=?
	`
	err = tx.Get(&generatedModel, enrollmentGeneratedQuery, generatedID)
	if err != nil {
		_ = tx.Rollback()
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	_ = tx.Commit()

	generated := fromModelToEnrollmentGenerated(&generatedModel)

	return &generated, nil
}
