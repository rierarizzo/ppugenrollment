package user

import (
	"ppugenrollment/internal/domain"
)

func fromStudentToStudentModel(student *domain.Student) StudentModel {
	return StudentModel{
		CommonFieldsModel: CommonFieldsModel{
			ID:           student.User.ID,
			IDCardNumber: student.User.IDCardNumber,
			Name:         student.User.Name,
			Surname:      student.User.Surname,
			Email:        student.User.Email,
			Password:     student.User.Password,
			Role:         student.User.Role,
		},
		DateOfBirth: student.DateOfBirth,
		IsAGraduate: student.IsAGraduate,
		Level:       student.Level,
	}
}

func fromAdminToAdminModel(admin *domain.Admin) AdminModel {
	return AdminModel{
		CommonFieldsModel{
			ID:           admin.User.ID,
			IDCardNumber: admin.User.IDCardNumber,
			Name:         admin.User.Name,
			Surname:      admin.User.Surname,
			Email:        admin.User.Email,
			Password:     admin.User.Password,
			Role:         admin.User.Role,
		}}
}

func fromApproverToApproverModel(approver *domain.Approver) ApproverModel {
	return ApproverModel{
		CommonFieldsModel{
			ID:           approver.User.ID,
			IDCardNumber: approver.User.IDCardNumber,
			Name:         approver.User.Name,
			Surname:      approver.User.Surname,
			Email:        approver.User.Email,
			Password:     approver.User.Password,
			Role:         approver.User.Role,
		}}
}

func fromCommonFieldsModelToUser(fields *CommonFieldsModel) domain.User {
	return domain.User{
		ID:           fields.ID,
		IDCardNumber: fields.IDCardNumber,
		Name:         fields.Name,
		Surname:      fields.Surname,
		Email:        fields.Email,
		Password:     fields.Password,
		Role:         fields.Role,
	}
}
