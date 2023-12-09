package authentication

import (
	"net/http"
	"ppugenrollment/internal/domain"
	auth "ppugenrollment/internal/usecases/user_authenticator"

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
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		user := fromRequestToUser(&request)
		appErr := userAuth.Register(&user)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		return c.JSON(http.StatusAccepted, "OK")
	}
}

func login(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request UserRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		authPayload, appErr := userAuth.Login(request.Email, request.Password)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		return c.JSON(http.StatusAccepted, fromAuthPayloadToResponse(authPayload))
	}
}
