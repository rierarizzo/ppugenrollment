package user_authenticator

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

func (a *DefaultAuthenticator) Register(userRegistrable domain.UserRegistrable) *domain.AppError {
	user := userRegistrable.GetUser()

	appErr := a.cryptUserPassword(user)
	if appErr != nil {
		return appErr
	}

	var repositoryErr *domain.AppError
	switch u := userRegistrable.(type) {
	case *domain.Student:
		repositoryErr = a.userRepo.InsertStudent(u)
	case *domain.Approver:
		repositoryErr = a.userRepo.InsertApprover(u)
	case *domain.Admin:
		repositoryErr = a.userRepo.InsertAdmin(u)
	default:
		slog.Error("Error while determining user type")
		repositoryErr = domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	if repositoryErr != nil {
		slog.Error(repositoryErr.Error())
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return nil
}

func (a *DefaultAuthenticator) Login(email, password string) (*domain.AuthUserPayload, *domain.AppError) {
	user, appErr := a.userRepo.SelectUserByEmail(email)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		slog.Error(err.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	token, appErr := security.CreateJWTToken(*user)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.NotAuthenticatedError)
	}

	payload := &domain.AuthUserPayload{
		User:        *user,
		AccessToken: token,
	}

	return payload, nil
}

func (a *DefaultAuthenticator) cryptUserPassword(user *domain.User) *domain.AppError {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.UnexpectedError)
	}

	user.Password = string(bytes)

	return nil
}
