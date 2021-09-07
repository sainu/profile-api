package models

// SocialLink is struct of social link
type SocialLink struct {
	Name string
	URL  string
}

// GetAllSocialLinks returns all social links
func GetAllSocialLinks() *[]SocialLink {
	return &[]SocialLink{
		{
			Name: "GitHub",
			URL:  "https://github.com/sainu",
		},
		{
			Name: "Twitter",
			URL:  "https://twitter.com/sainuio",
		},
		{
			Name: "Facebook",
			URL:  "https://www.facebook.com/sainou.katsutoshi",
		},
		{
			Name: "Wantedly",
			URL:  "https://www.wantedly.com/id/sainu",
		},
		{
			Name: "Qiita",
			URL:  "https://qiita.com/sainu",
		},
	}
}
