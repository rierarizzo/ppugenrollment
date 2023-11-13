package student_authenticator

import (
	"ppugenrollment/internal/domain"
)

type DefaultAuthenticator struct {
	userRepository UserRepository
}

func New(userRepository UserRepository) *DefaultAuthenticator {
	return &DefaultAuthenticator{userRepository: userRepository}
}

func (d DefaultAuthenticator) Register(student *domain.Student) *domain.AppError {
	return d.userRepository.InsertStudent(student)
}

func (d DefaultAuthenticator) Login(email, password string) (*domain.AuthenticatedUser, *domain.AppError) {
	student, appErr := d.userRepository.GetStudentByEmail(email)
	if appErr != nil {
		return nil, appErr
	}

	authenticatedUser := &domain.AuthenticatedUser{
		User:        student.User,
		AccessToken: "",
	}

	return authenticatedUser, nil
}
