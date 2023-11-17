package authenticator

import (
	"ppugenrollment/internal/domain"
)

type Authenticator interface {
	Register(user domain.User) *domain.AppError
	Login(email, password string) (*domain.AuthUserPayload, *domain.AppError)
}

type UserRepository interface {
	InsertStudent(student *domain.Student) *domain.AppError
	InsertAdmin(admin *domain.Admin) *domain.AppError
	InsertApprover(approver *domain.Approver) *domain.AppError

	SelectUserByEmail(email string) (*domain.CommonUserFields, *domain.AppError)
}
