package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"ppugenrollment/internal/api/mappers"
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/domain"
	"ppugenrollment/pkg/utils"
	"strconv"
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

func (pc *ProjectController) GetProjectByID(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	project, appErr := pc.manager.GetProjectByID(projectID)

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	response := mappers.FromProjectToResponse(project)

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

func (pc *ProjectController) UpdateProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	request := new(types.ProjectRequest)

	if err := c.Bind(&request); err != nil {
		slog.Error(err.Error())
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	project := mappers.FromRequestToProject(request)

	appErr := pc.manager.UpdateProject(projectID, &project)

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	return utils.SendOK(c, http.StatusAccepted, fmt.Sprintf("Project with ID %v updated", projectID), nil)
}

func (pc *ProjectController) DeleteProject(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))

	if err != nil {
		return utils.SendError(http.StatusBadRequest, domain.NewAppErrorWithType(domain.BadRequestError))
	}

	appErr := pc.manager.DeleteProject(projectID)

	if appErr != nil {
		return utils.SendError(http.StatusInternalServerError, appErr)
	}

	return utils.SendOK(c, http.StatusOK, fmt.Sprintf("Project with ID %v deleted", projectID), nil)
}
