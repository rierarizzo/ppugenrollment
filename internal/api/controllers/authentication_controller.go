package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
	auth "ppugenrollment/internal/ports"
)

func AuthRoutes(g *echo.Group) func(userAuth auth.Authenticator) {
	return func(userAuth auth.Authenticator) {
		g.POST("/register", register(userAuth))
		g.POST("/login", login(userAuth))
	}
}

// register is a function that handles the registration of a user.
// It takes an authenticator object as input and returns an echo.HandlerFunc.
// The authenticator object is used to register the user and returns an app error if the registration fails.
// The function binds the request to a UserRegisterRequest object and converts it to a User object.
// It then calls the Register method of the authenticator to register the user.
// If there is an app error, it returns a JSON response with the app error.
// Otherwise, it returns a JSON response with the status "OK".
func register(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request types.UserRegisterRequest
		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		validate := validator.New()
		err := validate.Struct(request)
		if err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		user := mappers.FromRegisterRequestToUser(&request)
		appErr := userAuth.Register(&user)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		return c.JSON(http.StatusAccepted, "OK")
	}
}

// login is a function that handles user login.
// It takes an authenticator object as input and returns an echo.HandlerFunc.
// The function binds the request to a UserRegisterRequest object.
// It then calls the Login method of the authenticator to authenticate the user.
// If there is an app error, it returns a JSON response with the app error.
// Otherwise, it returns a JSON response with the authenticated user's information in a UserResponse object.
func login(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request types.UserLoginRequest
		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		validate := validator.New()
		err := validate.Struct(request)
		if err != nil {
			slog.Error(err.Error())
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		authPayload, appErr := userAuth.Login(request.Email, request.Password)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		return c.JSON(http.StatusAccepted, mappers.FromAuthPayloadToResponse(authPayload))
	}
}
