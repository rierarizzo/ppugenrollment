package types

import (
	"ppugenrollment/pkg/domain"
)

type EnrollmentApplicationRequest struct {
	Project      int    `json:"project"`
	Schedule     int    `json:"schedule"`
	ScheduleCode string `json:"schedule_code"`
}

func (r *EnrollmentApplicationRequest) Validate() *domain.AppError {
	v := new(Validator)

	v.MustNotBeZero(r.Project)
	v.MustNotBeZero(r.Schedule)

	return v.AppErr
}

type EnrollmentApplicationResponse struct {
	ID             int    `json:"id,omitempty"`
	Student        int    `json:"student,omitempty"`
	StudentName    string `json:"student_name,omitempty"`
	StudentSurname string `json:"student_surname,omitempty"`
	Project        int    `json:"project,omitempty"`
	ProjectName    string `json:"project_name,omitempty"`
	CompanyName    string `json:"company_name,omitempty"`
	Schedule       int    `json:"schedule,omitempty"`
	ScheduleCode   string `json:"schedule_code,omitempty"`
	Status         string `json:"status,omitempty"`
}
