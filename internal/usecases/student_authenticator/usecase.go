package student_authenticator

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/security"
)

type DefaultAuthenticator struct {
	userRepo UserRepository
}

func New(userRepo UserRepository) *DefaultAuthenticator {
	return &DefaultAuthenticator{userRepo}
}

func (d DefaultAuthenticator) Register(student *domain.Student) *domain.AppError {
	bytes, err := bcrypt.GenerateFromPassword([]byte(student.User.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.UnexpectedError)
	}

	student.User.Password = string(bytes)

	appErr := d.userRepo.InsertStudent(student)
	if appErr != nil {
		slog.Error(err.Error())
		return domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return nil
}

func (d DefaultAuthenticator) Login(email, password string) (*domain.AuthenticatedUser, *domain.AppError) {
	notAuthErr := domain.NewAppErrorWithType(domain.NotAuthenticatedError)

	student, appErr := d.userRepo.GetStudentByEmail(email)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, notAuthErr
	}

	err := bcrypt.CompareHashAndPassword([]byte(student.User.Password), []byte(password))
	if err != nil {
		slog.Error(err.Error())
		return nil, notAuthErr
	}

	token, appErr := security.CreateJWTToken(student.User)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, notAuthErr
	}

	authenticatedUser := &domain.AuthenticatedUser{
		User:        student.User,
		AccessToken: token,
	}

	return authenticatedUser, nil
}
