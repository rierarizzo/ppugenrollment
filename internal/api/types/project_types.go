package types

import "time"

type ProjectResponse struct {
	ID          int             `json:"id"`
	Company     CompanyResponse `json:"company"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Schedule    string          `json:"schedule"`
	Starts      time.Time       `json:"starts"`
	Ends        time.Time       `json:"ends"`
}
