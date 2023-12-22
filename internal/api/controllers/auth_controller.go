package controllers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/api/utils"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
)

type AuthController struct {
	userAuth ports.Authenticator
}

func NewAuthController(userAuth ports.Authenticator) *AuthController {
	return &AuthController{userAuth}
}

func (ac *AuthController) Register(c echo.Context) error {
	request := new(types.UserRegisterRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return utils.SendError(http.StatusBadRequest, appErr)
	}

	user := mappers.FromRegisterRequestToUser(request)
	appErr := ac.userAuth.Register(&user)

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	return utils.SendOK(c, http.StatusCreated, "User registered", nil)
}

func (ac *AuthController) Login(c echo.Context) error {
	request := new(types.UserLoginRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return utils.SendError(http.StatusBadRequest, appErr)
	}

	authPayload, appErr := ac.userAuth.Login(request.Email, request.Password)

	if appErr != nil {
		return utils.SendError(http.StatusUnauthorized, appErr)
	}

	response := mappers.FromAuthPayloadToResponse(authPayload)

	return utils.SendOK(c, http.StatusAccepted, "User logged", response)
}
