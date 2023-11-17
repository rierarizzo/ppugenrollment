package authentication

import (
	"github.com/labstack/echo/v4"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/usecases/authenticator"
)

func Routes(g *echo.Group) func(studentAuthenticator authenticator.StudentAuthenticator,
	adminAuthenticator authenticator.AdminAuthenticator, approverAuthenticator authenticator.ApproverAuthenticator) {
	return func(studentAuth authenticator.StudentAuthenticator, adminAuth authenticator.AdminAuthenticator,
		approverAuth authenticator.ApproverAuthenticator) {
		g.POST("/student/register", registerStudent(studentAuth))
		g.POST("/admin/register", registerAdmin(adminAuth))
		g.POST("/approver/register", registerApprover(approverAuth))
	}
}

func registerStudent(studentAuthenticator authenticator.StudentAuthenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request StudentRequest
		if err := c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		student := fromRequestToStudent(&request)

		appErr := studentAuthenticator.Register(&student)
		if appErr != nil {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil
	}
}

func registerAdmin(adminAuthenticator authenticator.AdminAuthenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request AdminRequest
		if err := c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		admin := fromRequestToAdmin(&request)

		appErr := adminAuthenticator.Register(&admin)
		if appErr != nil {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil
	}
}

func registerApprover(approverAuthenticator authenticator.ApproverAuthenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ApproverRequest
		if err := c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		approver := fromRequestToApprover(&request)

		appErr := approverAuthenticator.Register(&approver)
		if appErr != nil {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil
	}
}
