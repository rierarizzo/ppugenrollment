package controllers

import (
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

// register is a function that registers a new user
func register(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(types.UserRegisterRequest)

		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
		}

		if appErr := request.Validate(); appErr != nil {
			slog.Error(appErr.Error())
			return sendError(http.StatusBadRequest, appErr)
		}

		user := mappers.FromRegisterRequestToUser(request)
		appErr := userAuth.Register(&user)

		if appErr != nil {
			return sendError(http.StatusInternalServerError, appErr)
		}

		return sendOK(c, http.StatusCreated, "User registered", nil)
	}
}

// login is a function that logs in a user
func login(userAuth auth.Authenticator) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := new(types.UserLoginRequest)

		if err := c.Bind(&request); err != nil {
			slog.Error(err.Error())
			return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
		}

		if appErr := request.Validate(); appErr != nil {
			slog.Error(appErr.Error())
			return sendError(http.StatusBadRequest, appErr)
		}

		authPayload, appErr := userAuth.Login(request.Email, request.Password)

		if appErr != nil {
			return sendError(http.StatusUnauthorized, appErr)
		}

		response := mappers.FromAuthPayloadToResponse(authPayload)

		return sendOK(c, http.StatusAccepted, "User logged", response)
	}
}
