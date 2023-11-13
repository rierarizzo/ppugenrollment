package authentication

import (
	"ppugenrollment/internal/domain"
)

func fromRequestToStudent(request *StudentRequest) domain.Student {
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
