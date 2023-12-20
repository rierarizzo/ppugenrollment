package types

import "time"

type EnrollmentGeneratedResponse struct {
	ID                    int                `json:"id"`
	EnrollmentApplication int                `json:"application_id"`
	Project               ProjectResponse    `json:"project"`
	ApprovedBy            ApprovedByResponse `json:"approvedBy"`
	GeneratedAt           time.Time          `json:"generated_at"`
}

type CompanyResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	RUC  string `json:"ruc"`
}

type ApprovedByResponse struct {
	ID           int    `json:"id"`
	IDCardNumber string `json:"card_number"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
}
