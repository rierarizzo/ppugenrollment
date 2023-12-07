package approval

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/usecases/enrollment_application_approver"
	"strconv"
)

func Routes(g *echo.Group) func(approver enrollment_application_approver.Approver) {
	return func(approver enrollment_application_approver.Approver) {
		g.POST("/approve/:application_id", approveEnrollmentApplication(approver))
	}
}

func approveEnrollmentApplication(approver enrollment_application_approver.Approver) echo.HandlerFunc {
	return func(c echo.Context) error {
		applicationID, err := strconv.Atoi(c.Param("application_id"))
		if err != nil {
			return domain.NewAppError(err, domain.BadRequestError)
		}

		generated, appErr := approver.ApproveEnrollmentApplication(applicationID, 1)
		if appErr != nil {
			return appErr
		}

		response := fromGeneratedToResponse(generated)

		return c.JSON(http.StatusAccepted, response)
	}
}
