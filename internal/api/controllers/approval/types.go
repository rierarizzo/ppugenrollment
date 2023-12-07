package approval

import "time"

type EnrollmentGeneratedResponse struct {
	ID                    int                `json:"id"`
	EnrollmentApplication int                `json:"application_id"`
	Project               projectResponse    `json:"project"`
	ApprovedBy            approvedByResponse `json:"approvedBy"`
	GeneratedAt           time.Time          `json:"generated_at"`
}

type projectResponse struct {
	ID          int             `json:"id"`
	Company     companyResponse `json:"company"`
	Description string          `json:"description"`
	Schedule    string          `json:"schedule"`
	Starts      time.Time       `json:"starts"`
	Ends        time.Time       `json:"ends"`
}

type companyResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	RUC  string `json:"ruc"`
}

type approvedByResponse struct {
	ID           int    `json:"id"`
	IDCardNumber string `json:"card_number"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
}
