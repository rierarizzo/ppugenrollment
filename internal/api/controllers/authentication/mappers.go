package authentication

import (
	"ppugenrollment/internal/domain"
)

func fromRequestToStudent(request *StudentRequest) domain.Student {
	return domain.Student{
		UserFields: domain.CommonUserFields{
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

func fromRequestToAdmin(request *AdminRequest) domain.Admin {
	return domain.Admin{
		UserFields: domain.CommonUserFields{
			IDCardNumber: request.IDCardNumber,
			Name:         request.Name,
			Surname:      request.Surname,
			Email:        request.Email,
			Password:     request.Password,
			Role:         request.Role,
		}}
}

func fromRequestToApprover(request *ApproverRequest) domain.Approver {
	return domain.Approver{
		UserFields: domain.CommonUserFields{
			IDCardNumber: request.IDCardNumber,
			Name:         request.Name,
			Surname:      request.Surname,
			Email:        request.Email,
			Password:     request.Password,
			Role:         request.Role,
		}}
}
