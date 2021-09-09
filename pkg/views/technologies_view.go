package views

import "github.com/sainu/profile-api/pkg/models"

// TechnologyView is struct of technology view
type TechnologyView struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

// NewTechnologiesView is constructor for view of technologies
func NewTechnologiesView(technologies *[]models.Technology) *[]TechnologyView {
	var view []TechnologyView

	for _, technology := range *technologies {
		view = append(view, TechnologyView{
			Name:     technology.Name,
			Versions: technology.Versions,
		})
	}

	return &view
}
