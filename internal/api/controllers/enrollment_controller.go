package controllers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
)

type EnrollmentController struct {
	enroller ports.ProjectEnroller
}

func NewEnrollmentController(enroller ports.ProjectEnroller) *EnrollmentController {
	return &EnrollmentController{enroller}
}

func (ec *EnrollmentController) EnrollToProject(c echo.Context) error {
	request := new(types.EnrollmentApplicationRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	enrolledBy := c.Get("UserID").(int)

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return sendError(http.StatusBadRequest, appErr)
	}

	enrollmentApplication := mappers.FromRequestToApplication(request)
	application, appErr := ec.enroller.EnrollToProject(&enrollmentApplication, enrolledBy)

	if appErr != nil {
		return sendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromApplicationToResponse(application)

	return sendOK(c, http.StatusAccepted, "Enrollment applied", response)
}
