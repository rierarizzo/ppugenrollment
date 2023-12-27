package controllers

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/domain"
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
		return sendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromProjectsToResponse(projects)

	return sendOK(c, http.StatusOK, "", response)
}

func (pc *ProjectController) AddNewProject(c echo.Context) error {
	request := new(types.ProjectRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return sendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	if appErr := request.Validate(); appErr != nil {
		slog.Error(appErr.Error())
		return sendError(http.StatusBadRequest, appErr)
	}

	project := mappers.FromRequestToProject(request)

	projectWithID, appErr := pc.manager.AddNewProject(&project)

	if appErr != nil {
		return sendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromProjectToResponse(projectWithID)

	return sendOK(c, http.StatusAccepted, "New project created", response)
}
