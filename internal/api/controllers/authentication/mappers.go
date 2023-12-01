package authentication

import (
	"ppugenrollment/internal/domain"
)

func fromRequestToStudent(request *UserRequest) domain.Student {
	return domain.Student{
		User: domain.User{
			IDCardNumber: request.IDCardNumber,
			Name:         request.Name,
			Surname:      request.Surname,
			Email:        request.Email,
			Password:     request.Password,
			Role:         request.Role,
		},
		DateOfBirth: request.DateOfBirth,
		IsAGraduate: request.IsAGraduate,
		Level:       request.Level,
	}
}

func fromRequestToAdmin(request *UserRequest) domain.Admin {
	return domain.Admin{
		User: domain.User{
			IDCardNumber: request.IDCardNumber,
			Name:         request.Name,
			Surname:      request.Surname,
			Email:        request.Email,
			Password:     request.Password,
			Role:         request.Role,
		}}
}

func fromRequestToApprover(request *UserRequest) domain.Approver {
	return domain.Approver{
		User: domain.User{
			IDCardNumber: request.IDCardNumber,
			Name:         request.Name,
			Surname:      request.Surname,
			Email:        request.Email,
			Password:     request.Password,
			Role:         request.Role,
		}}
}

func fromAuthPayloadToResponse(payload *domain.AuthUserPayload) UserResponse {
	return UserResponse{
		ID:      payload.ID,
		Name:    payload.Name,
		Surname: payload.Surname,
		Email:   payload.Email,
		Role:    payload.Role,
		AuthenticationResult: struct {
			AccessToken string `json:"access_token"`
		}(struct{ AccessToken string }{
			AccessToken: payload.AccessToken,
		}),
	}
}
