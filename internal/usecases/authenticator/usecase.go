package authenticator

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/security"
)

type UserAuthenticator struct {
	userRepo UserRepository
}

func New(userRepo UserRepository) *UserAuthenticator {
	return &UserAuthenticator{userRepo}
}

func (a *UserAuthenticator) Register(userRegistrable domain.UserRegistrable) *domain.AppError {
	user := userRegistrable.GetUser()

	appErr := a.cryptUserPassword(user)
	if appErr != nil {
		return appErr
	}

	switch u := userRegistrable.(type) {
	case *domain.Student:
		return a.userRepo.InsertStudent(u)
	case *domain.Approver:
		return a.userRepo.InsertApprover(u)
	case *domain.Admin:
		return a.userRepo.InsertAdmin(u)
	default:
		slog.Error("Error while determining user type")
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}
}

func (a *UserAuthenticator) Login(email, password string) (*domain.AuthUserPayload, *domain.AppError) {
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

func (a *UserAuthenticator) cryptUserPassword(user *domain.User) *domain.AppError {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.UnexpectedError)
	}

	user.Password = string(bytes)

	return nil
}
