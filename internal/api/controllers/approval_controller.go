package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/utils"
	"strconv"
)

type ApprovalController struct {
	approver ports.EnrollmentApprover
}

func NewApprovalController(approver ports.EnrollmentApprover) *ApprovalController {
	return &ApprovalController{approver}
}

func (ac *ApprovalController) ApproveEnrollmentApplication(c echo.Context) error {
	applicationID, err := strconv.Atoi(c.Param("application_id"))

	if err != nil {
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	approvedBy := c.Get("UserID").(int)

	generated, appErr := ac.approver.ApproveEnrollmentApplication(applicationID, approvedBy)

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromGeneratedToResponse(generated)

	return utils.SendOK(c, http.StatusAccepted, "Enrollment application approved", response)

}
