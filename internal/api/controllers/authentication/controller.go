package authentication

import (
	"ppugenrollment/internal/domain"
	auth "ppugenrollment/internal/usecases/authenticator"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) func(
	studentAuth auth.StudentAuthenticator, adminAuth auth.AdminAuthenticator, approverAuth auth.ApproverAuthenticator,
) {
	return func(
		studentAuth auth.StudentAuthenticator,
		adminAuth auth.AdminAuthenticator,
		approverAuth auth.ApproverAuthenticator,
	) {
		g.POST("/student/register", registerStudent(studentAuth))
		g.POST("/admin/register", registerAdmin(adminAuth))
		g.POST("/approver/register", registerApprover(approverAuth))
	}
}

func registerStudent(studentAuthenticator auth.StudentAuthenticator) echo.HandlerFunc {
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

func registerAdmin(adminAuthenticator auth.AdminAuthenticator) echo.HandlerFunc {
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

func registerApprover(approverAuthenticator auth.ApproverAuthenticator) echo.HandlerFunc {
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
