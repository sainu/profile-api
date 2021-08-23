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
			Name:  "TypeScript",
			Score: 50,
		},
		{
			Name:  "Ruby",
			Score: 60,
		},
		{
			Name:  "Go",
			Score: 40,
		},
		{
			Name:  "Ruby on Rails",
			Score: 67,
		},
		{
			Name:  "Vue.js",
			Score: 58,
		},
		{
			Name:  "Nuxt.js",
			Score: 55,
		},
		{
			Name:  "React.js",
			Score: 48,
		},
		{
			Name:  "Next.js",
			Score: 45,
		},
		{
			Name:  "MySQL",
			Score: 55,
		},
		{
			Name:  "Redis",
			Score: 55,
		},
		{
			Name:  "Docker",
			Score: 53,
		},
		{
			Name:  "S3",
			Score: 53,
		},
		{
			Name:  "EC2",
			Score: 53,
		},
	}
}
