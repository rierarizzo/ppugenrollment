package models

import "time"

type EnrollmentGeneratedModel struct {
	ID                    int       `db:"id"`
	EnrollmentApplication int       `db:"application_id"`
	ProjectID             int       `db:"project_id"`
	CompanyID             int       `db:"company_id"`
	CompanyName           string    `db:"company_name"`
	CompanyRUC            string    `db:"company_ruc"`
	ProjectDescription    string    `db:"project_description"`
	ProjectSchedule       string    `db:"project_schedule"`
	ProjectStarts         time.Time `db:"project_starts"`
	ProjectEnds           time.Time `db:"project_ends"`
	ApproverID            int       `db:"approver_id"`
	ApproverIDCardNumber  string    `db:"approver_card_number"`
	ApproverName          string    `db:"approver_name"`
	ApproverSurname       string    `db:"approver_surname"`
	GeneratedAt           time.Time `db:"generated_at"`
}
