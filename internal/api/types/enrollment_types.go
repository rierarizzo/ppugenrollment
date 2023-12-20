package types

type EnrollmentApplicationRequest struct {
	Student  int `json:"student,omitempty"`
	Project  int `json:"project"`
	Schedule int `json:"schedule"`
}

type EnrollmentApplicationResponse struct {
	ID       int `json:"id,omitempty"`
	Student  int `json:"student,omitempty"`
	Project  int `json:"project,omitempty"`
	Schedule int `json:"schedule,omitempty"`
}
