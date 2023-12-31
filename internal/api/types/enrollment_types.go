package types

import (
	"ppugenrollment/pkg/domain"
)

type EnrollmentApplicationRequest struct {
	Project  int `json:"project"`
	Schedule int `json:"schedule"`
}

func (r *EnrollmentApplicationRequest) Validate() *domain.AppError {
	v := new(Validator)

	v.MustNotBeZero(r.Project)
	v.MustNotBeZero(r.Schedule)

	return v.AppErr
}

type EnrollmentApplicationResponse struct {
	ID       int `json:"id,omitempty"`
	Student  int `json:"student,omitempty"`
	Project  int `json:"project,omitempty"`
	Schedule int `json:"schedule,omitempty"`
}
