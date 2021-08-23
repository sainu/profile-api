package views

import (
	"github.com/sainu/profile-api/models"
)

// SkillView is struct of skill view
type SkillView struct {
	Name  string `json:"name"`
	Score uint8  `json:"score"`
}

// NewSkillsView is constructor for view of skills
func NewSkillsView(skills []models.Skill) []SkillView {
	var view []SkillView

	for _, skill := range skills {
		view = append(view, SkillView{
			Name:  skill.Name,
			Score: skill.Score,
		})
	}

	return view
}
