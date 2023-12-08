package enrollment

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/usecases/project_enroller"
)

func Routes(g *echo.Group) func(enroller project_enroller.Enroller) {
	return func(enroller project_enroller.Enroller) {
		g.POST("/enrollToProject", enrollToProject(enroller))
	}
}

func enrollToProject(enroller project_enroller.Enroller) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request ApplicationRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, domain.NewAppError(err, domain.BadRequestError))
		}

		enrollmentApplication := fromRequestToApplication(&request)
		application, appErr := enroller.EnrollToProject(&enrollmentApplication)
		if appErr != nil {
			return c.JSON(http.StatusInternalServerError, appErr)
		}

		response := fromApplicationToResponse(application)

		return c.JSON(http.StatusAccepted, response)
	}
}
