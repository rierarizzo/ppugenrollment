package authenticator

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
)

type AdminAuthenticator struct {
	userRepo UserRepository
}

func NewAdminAuthenticator(userRepo UserRepository) *AdminAuthenticator {
	return &AdminAuthenticator{userRepo}
}

func (a *AdminAuthenticator) Register(user domain.User) *domain.AppError {
	var admin *domain.Admin
	if a, ok := user.(*domain.Admin); ok {
		admin = a
	} else {
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(admin.UserFields.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.UnexpectedError)
	}

	admin.UserFields.Password = string(bytes)

	appErr := a.userRepo.InsertAdmin(admin)
	if appErr != nil {
		slog.Error(appErr.Error())
		return domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return nil
}

func (a *AdminAuthenticator) Login(email, password string) (*domain.AuthUserPayload, *domain.AppError) {
	return loginInSystem(email, password, a.userRepo)
}
