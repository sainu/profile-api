package models

import (
	"net/http"
	"strings"
	"time"

	"github.com/sainu/profile-api/pkg/microcms"
)

// Experience is struct of experience
type Experience struct {
	CompanyName    string
	EmploymentType string
	StartDate      time.Time
	EndDate        time.Time
	Projects       []Project
}

const (
	// FullTimeEmployment is "正社員"
	FullTimeEmployment = "正社員"
	// PartTimeEmployment is "パートタイム"
	PartTimeEmployment = "パートタイム"
	// Freelance is "フリーランス"
	Freelance = "フリーランス"
	// Internship is "インターン"
	Internship = "インターン"
)

// GetAllExperiences returns all expericnes
func GetAllExperiences() *[]Experience {
	client := microcms.NewClient()
	req := microcms.NewRequest(
		http.MethodGet,
		"/api/v1/experiences",
		microcms.WithQuery(map[string]string{"limit": "20", "orders": "-startDate"}),
	)
	resp := new(microcms.ExperienceResponseBody)
	client.Do(req, resp)

	var experiences []Experience
	for _, c := range resp.Contents {

		var projects []Project
		for _, p := range c.Projects {

			var technologies []Technology
			for _, t := range p.Technologies {
				technologies = append(technologies, Technology{
					Name:     t.Technology.Name,
					Versions: strings.Split(t.Versions, ","),
				})
			}

			projects = append(projects, Project{
				Description:  p.Description,
				Technologies: technologies,
			})
		}

		startDate, _ := time.Parse(time.RFC3339, c.StartDate)
		endDate, _ := time.Parse(time.RFC3339, c.EndDate)
		experiences = append(experiences, Experience{
			CompanyName:    c.CompanyName,
			EmploymentType: c.EmploymentTypes[0],
			StartDate:      startDate,
			EndDate:        endDate,
			Projects:       projects,
		})
	}

	return &experiences
}
