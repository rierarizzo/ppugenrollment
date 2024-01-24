package controllers

import (
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/domain"
	"ppugenrollment/pkg/utils"

	"github.com/labstack/echo/v4"
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
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	enrolledBy := c.Get("UserID").(int)

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return utils.SendError(http.StatusBadRequest, appErr)
	}

	enrollmentApplication := mappers.FromRequestToApplication(request)
	application, appErr := ec.enroller.EnrollToProject(&enrollmentApplication, enrolledBy)

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromApplicationToResponse(application)

	return utils.SendOK(c, http.StatusAccepted, "Enrollment applied", response)
}

func (ec *EnrollmentController) GetEnrollmentApplications(c echo.Context) error {
	applications, appErr := ec.enroller.GetEnrollmentApplications()

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromApplicationsToResponse(applications)

	return utils.SendOK(c, http.StatusOK, "Enrollments retrieved", response)
}
