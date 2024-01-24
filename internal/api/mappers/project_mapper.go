package mappers

import (
	"ppugenrollment/internal/api/types"
	"ppugenrollment/pkg/domain"
)

func FromRequestToProject(request *types.ProjectRequest) domain.Project {
	return domain.Project{
		Company:     domain.Company{ID: request.Company},
		Name:        request.Name,
		Description: request.Description,
		Schedules:   request.Schedules,
		Starts:      request.Starts,
		Ends:        request.Ends,
	}
}

func FromProjectToResponse(project *domain.Project) types.ProjectResponse {
	return types.ProjectResponse{
		ID: project.ID,
		Company: types.CompanyResponse{
			ID:   project.Company.ID,
			Name: project.Company.Name,
			RUC:  project.Company.RUC,
		},
		Name:        project.Name,
		Description: project.Description,
		Schedules:   project.Schedules,
		Starts:      project.Starts,
		Ends:        project.Ends,
	}
}

func FromCompaniesToResponse(companies []domain.Company) []types.CompanyResponse {
	var response []types.CompanyResponse
	for _, v := range companies {
		response = append(response, types.CompanyResponse{
			ID:   v.ID,
			Name: v.Name,
			RUC:  v.RUC,
		})
	}

	return response
}

func FromProjectsToResponse(projects []domain.Project) []types.ProjectResponse {
	var response []types.ProjectResponse
	for _, v := range projects {
		response = append(response, FromProjectToResponse(&v))
	}

	return response
}
