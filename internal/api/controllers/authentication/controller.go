package authentication

import (
	"net/http"
	"ppugenrollment/internal/domain"
	auth "ppugenrollment/internal/usecases/authenticator"

	"github.com/labstack/echo/v4"
)

func Routes(g *echo.Group) func(userAuth auth.Authenticator) {
	return func(userAuth auth.Authenticator) {
		g.POST("/register", register(userAuth))
		g.POST("/login", login(userAuth))
	}
}

func register(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request UserRequest
		if err := c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		var authenticatorErr *domain.AppError
		switch request.Role {
		case "S": // Student
			student := fromRequestToStudent(&request)
			authenticatorErr = userAuth.Register(&student)
			break
		case "A": // Approver
			approver := fromRequestToApprover(&request)
			authenticatorErr = userAuth.Register(&approver)
			break
		case "M": // Admin
			admin := fromRequestToAdmin(&request)
			authenticatorErr = userAuth.Register(&admin)
		default:
			return domain.NewAppErrorWithType(domain.BadRequestError)
		}

		if authenticatorErr != nil {
			return domain.NewAppError(authenticatorErr, domain.UnexpectedError)
		}

		return nil
	}
}

func login(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request UserRequest
		if err := c.Bind(&request); err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		authPayload, appErr := userAuth.Login(request.Email, request.Password)
		if appErr != nil {
			return domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return c.JSON(http.StatusAccepted, fromAuthPayloadToResponse(authPayload))
	}
}
