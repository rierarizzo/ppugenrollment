package student

import (
	"ppugenrollment/internal/domain"
)

func fromStudentToModel(student *domain.Student) Model {
	return Model{
		ID:           student.User.ID,
		IDCardNumber: student.User.IDCardNumber,
		Name:         student.User.Name,
		Surname:      student.User.Surname,
		Email:        student.User.Email,
		Password:     student.User.Password,
		Role:         student.User.Role,
		DateOfBirth:  student.DateOfBirth,
		IsAGraduate:  student.IsAGraduate,
		Level:        student.Level,
	}
}

func fromModelToStudent(model *Model) domain.Student {
	return domain.Student{
		User: domain.User{
			ID:           model.ID,
			IDCardNumber: model.IDCardNumber,
			Name:         model.Name,
			Surname:      model.Surname,
			Email:        model.Email,
			Password:     model.Password,
			Role:         model.Role,
		},
		DateOfBirth: model.DateOfBirth,
		IsAGraduate: model.IsAGraduate,
		Level:       model.Level,
	}
}
