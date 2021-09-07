package models

// Skill is struct of skill
type Skill struct {
	Name  string
	Score uint8
}

// GetAllSkills returns all skills
func GetAllSkills() *[]Skill {
	return &[]Skill{
		{
			Name:  "Ruby on Rails",
			Score: 85,
		},
		{
			Name:  "Ruby",
			Score: 82,
		},
		{
			Name:  "TypeScript",
			Score: 73,
		},
		{
			Name:  "Vue.js",
			Score: 70,
		},
		{
			Name:  "Nuxt.js",
			Score: 70,
		},
		{
			Name:  "Next.js",
			Score: 63,
		},
		{
			Name:  "React.js",
			Score: 60,
		},
		{
			Name:  "Docker",
			Score: 53,
		},
		{
			Name:  "MySQL",
			Score: 45,
		},
		{
			Name:  "Redis",
			Score: 40,
		},
		{
			Name:  "Go",
			Score: 20,
		},
	}
}
