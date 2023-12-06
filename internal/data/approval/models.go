package approval

import "time"

type EnrollmentGeneratedModel struct {
	ID                    int `db:"id"`
	EnrollmentApplication int `db:"application_id"`
	Project               struct {
		ID      int `db:"project_id"`
		Company struct {
			ID   int    `db:"company_id"`
			Name string `db:"company_name"`
			RUC  string `db:"company_ruc"`
		}
		Description string    `db:"project_description"`
		Schedule    string    `db:"project_schedule"`
		Starts      time.Time `db:"project_starts"`
		Ends        time.Time `db:"project_ends"`
	}
	ApprovedBy struct {
		ID           int    `db:"approver_id"`
		IDCardNumber string `db:"approver_card_number"`
		Name         string `db:"approver_name"`
		Surname      string `db:"approver_surname"`
	}
	GeneratedAt time.Time `db:"generated_at"`
}
