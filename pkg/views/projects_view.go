package views

import "github.com/sainu/profile-api/pkg/models"

// ProjectView is struct of project view
type ProjectView struct {
	Description  string            `json:"description"`
	Technologies *[]TechnologyView `json:"technologies"`
}

// NewProjectsView is constructor for view of project
func NewProjectsView(projects *[]models.Project) *[]ProjectView {
	var view []ProjectView

	for _, project := range *projects {
		view = append(view, ProjectView{
			Description:  project.Description,
			Technologies: NewTechnologiesView(&project.Technologies),
		})
	}

	return &view
}
