package mappers

import (
	"ppugenrollment/internal/data/models"
	"ppugenrollment/internal/domain"
)

func FromModelToUser(fields *models.UserModel) domain.User {
	return domain.User{
		ID:           fields.ID,
		IDCardNumber: fields.IDCardNumber,
		Name:         fields.Name,
		Surname:      fields.Surname,
		Email:        fields.Email,
		Password:     fields.Password,
		Role:         fields.Role,
		DateOfBirth:  fields.DateOfBirth,
		IsAGraduate:  fields.IsAGraduate,
		Level:        fields.Level,
	}
}

func FromUserToModel(user *domain.User) models.UserModel {
	return models.UserModel{
		ID:           user.ID,
		IDCardNumber: user.IDCardNumber,
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		Password:     user.Password,
		Role:         user.Role,
		DateOfBirth:  user.DateOfBirth,
		IsAGraduate:  user.IsAGraduate,
		Level:        user.Level,
	}
}
