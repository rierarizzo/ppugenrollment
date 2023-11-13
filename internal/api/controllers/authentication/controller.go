package authentication

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/usecases/student_authenticator"
)

func Routes(g *echo.Group) func(authenticator student_authenticator.Authenticator) {
	return func(a student_authenticator.Authenticator) {
		g.POST("/student/register", registerStudent(a))
	}
}

func registerStudent(authenticator student_authenticator.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request StudentRequest
		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return domain.NewAppError(err, domain.BadRequestError)
		}

		student := fromRequestToStudent(&request)

		appErr := authenticator.Register(&student)
		if appErr != nil {
			slog.Error(appErr.Error())
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil
	}
}
