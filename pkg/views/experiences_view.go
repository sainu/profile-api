package views

import (
	"time"

	"github.com/sainu/profile-api/pkg/models"
)

// ExperienceView is struct of experience view
type ExperienceView struct {
	CompanyName    string        `json:"company_name"`
	EmploymentType string        `json:"employment_type"`
	Department     string        `json:"department"`
	StartDate      string        `json:"start_date"`
	EndDate        string        `json:"end_date"`
	Projects       []ProjectView `json:"projects"`
}

// NewExperiencesView is constructor for view of experiences
func NewExperiencesView(experiences []models.Experience) []ExperienceView {
	var view []ExperienceView

	for _, experience := range experiences {
		view = append(view, ExperienceView{
			CompanyName:    experience.CompanyName,
			EmploymentType: experience.EmploymentType,
			Department:     experience.Department,
			StartDate:      experience.StartDate.Format(time.RFC3339),
			EndDate:        experience.EndDate.Format(time.RFC3339),
			Projects:       NewProjectsView(experience.Projects),
		})
	}

	return view
}
