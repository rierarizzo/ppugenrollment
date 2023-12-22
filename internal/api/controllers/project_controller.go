package controllers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/utils"
)

type ProjectController struct {
	manager ports.ProjectManager
}

func NewProjectController(manager ports.ProjectManager) *ProjectController {
	return &ProjectController{manager}
}

func (pc *ProjectController) GetAllProjects(c echo.Context) error {
	projects, appErr := pc.manager.GetAllProjects()

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromProjectsToResponse(projects)

	return utils.SendOK(c, http.StatusOK, "", response)
}

func (pc *ProjectController) AddNewProject(c echo.Context) error {
	request := new(types.ProjectRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return utils.SendError(http.StatusBadRequest, appErr)
	}

	project := mappers.FromRequestToProject(request)

	projectWithID, appErr := pc.manager.AddNewProject(&project)

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromProjectToResponse(projectWithID)

	return utils.SendOK(c, http.StatusAccepted, "New project created", response)
}
