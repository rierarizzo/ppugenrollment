package mappers

import (
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
)

func FromRequestToProject(request *types.ProjectRequest) domain.Project {
	return domain.Project{
		Company:     domain.Company{ID: request.Company},
		Name:        request.Name,
		Description: request.Description,
		Starts:      request.Starts,
		Ends:        request.Ends,
	}
}

func FromProjectToResponse(project *domain.Project) types.ProjectResponse {
	return types.ProjectResponse{
		ID:          project.ID,
		Company:     types.CompanyResponse{ID: project.Company.ID},
		Name:        project.Name,
		Description: project.Description,
		Starts:      project.Starts,
		Ends:        project.Ends,
	}
}

func FromProjectsToResponse(projects []domain.Project) []types.ProjectResponse {
	var response []types.ProjectResponse
	for _, v := range projects {
		response = append(response, FromProjectToResponse(&v))
	}

	return response
}
