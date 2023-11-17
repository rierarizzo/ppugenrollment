package user

import (
	"ppugenrollment/internal/domain"
)

func fromStudentToStudentModel(student *domain.Student) StudentModel {
	return StudentModel{
		CommonFieldsModel: CommonFieldsModel{
			ID:           student.UserFields.ID,
			IDCardNumber: student.UserFields.IDCardNumber,
			Name:         student.UserFields.Name,
			Surname:      student.UserFields.Surname,
			Email:        student.UserFields.Email,
			Password:     student.UserFields.Password,
			Role:         student.UserFields.Role,
		},
		DateOfBirth: student.DateOfBirth,
		IsAGraduate: student.IsAGraduate,
		Level:       student.Level,
	}
}

func fromAdminToAdminModel(admin *domain.Admin) AdminModel {
	return AdminModel{CommonFieldsModel{
		ID:           admin.UserFields.ID,
		IDCardNumber: admin.UserFields.IDCardNumber,
		Name:         admin.UserFields.Name,
		Surname:      admin.UserFields.Surname,
		Email:        admin.UserFields.Email,
		Password:     admin.UserFields.Password,
		Role:         admin.UserFields.Role,
	}}
}

func fromApproverToApproverModel(approver *domain.Approver) ApproverModel {
	return ApproverModel{CommonFieldsModel{
		ID:           approver.UserFields.ID,
		IDCardNumber: approver.UserFields.IDCardNumber,
		Name:         approver.UserFields.Name,
		Surname:      approver.UserFields.Surname,
		Email:        approver.UserFields.Email,
		Password:     approver.UserFields.Password,
		Role:         approver.UserFields.Role,
	}}
}

func fromCommonFieldsModelToUser(fields *CommonFieldsModel) domain.CommonUserFields {
	return domain.CommonUserFields{
		ID:           fields.ID,
		IDCardNumber: fields.IDCardNumber,
		Name:         fields.Name,
		Surname:      fields.Surname,
		Email:        fields.Email,
		Password:     fields.Password,
		Role:         fields.Role,
	}
}
