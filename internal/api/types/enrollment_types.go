package types

type EnrollmentApplicationRequest struct {
	Project  int `json:"project" validate:"required"`
	Schedule int `json:"schedule" validate:"required"`
}

type EnrollmentApplicationResponse struct {
	ID       int `json:"id,omitempty"`
	Student  int `json:"student,omitempty"`
	Project  int `json:"project,omitempty"`
	Schedule int `json:"schedule,omitempty"`
}
