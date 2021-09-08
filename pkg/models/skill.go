package models

import "github.com/sainu/profile-api/pkg/microcms"

// Skill is struct of skill
type Skill struct {
	Name  string
	Score uint8
}

// GetAllSkills returns all skills
func GetAllSkills() *[]Skill {
	client := microcms.NewClient()
	resp := new(microcms.SkillsResponseBody)
	client.Do(microcms.NewRequest("/api/v1/skills"), resp)

	var skills []Skill
	for _, c := range resp.Contents {
		skills = append(skills, Skill{
			Name:  c.Name,
			Score: c.Score,
		})
	}
	return &skills
}
