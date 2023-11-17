package authenticator

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"ppugenrollment/internal/domain"
)

type ApproverAuthenticator struct {
	userRepo UserRepository
}

func NewApproverAuthenticator(userRepo UserRepository) *ApproverAuthenticator {
	return &ApproverAuthenticator{userRepo}
}

func (a ApproverAuthenticator) Register(user domain.User) *domain.AppError {
	var approver *domain.Approver
	if a, ok := user.(*domain.Approver); ok {
		approver = a
	} else {
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(approver.UserFields.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.UnexpectedError)
	}

	approver.UserFields.Password = string(bytes)

	appErr := a.userRepo.InsertApprover(approver)
	if appErr != nil {
		slog.Error(appErr.Error())
		return domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return nil
}

func (a ApproverAuthenticator) Login(email, password string) (*domain.AuthUserPayload, *domain.AppError) {
	return loginInSystem(email, password, a.userRepo)
}
