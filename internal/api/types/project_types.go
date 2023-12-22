package types

import (
	"ppugenrollment/internal/domain"
	"ppugenrollment/pkg/utils"
	"time"
)

type ProjectRequest struct {
	Company     int       `json:"company"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Schedules   []string  `json:"schedules"`
	Starts      time.Time `json:"starts"`
	Ends        time.Time `json:"ends"`
}

func (r *ProjectRequest) Validate() *domain.AppError {
	v := new(utils.Validator)

	v.MustNotBeZero(r.Company)
	v.MustNotBeEmptyString(r.Name)
	v.MustNotBeEmptyString(r.Description)
	v.MustNotBeEmptyStringArray(r.Schedules)

	return v.AppErr
}

type ProjectResponse struct {
	ID          int             `json:"id"`
	Company     CompanyResponse `json:"company"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Schedule    string          `json:"schedule"`
	Starts      time.Time       `json:"starts"`
	Ends        time.Time       `json:"ends"`
}
