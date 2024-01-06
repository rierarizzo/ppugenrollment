package mappers

import (
	"ppugenrollment/internal/data/sqlcgen"
	"ppugenrollment/pkg/domain"
)

func FromModelToProject(model *sqlcgen.Project) domain.Project {
	return domain.Project{
		ID:          int(model.ID),
		Company:     domain.Company{ID: int(model.Company)},
		Name:        model.Name,
		Description: model.Description,
		Starts:      model.Starts,
		Ends:        model.Ends,
	}
}
