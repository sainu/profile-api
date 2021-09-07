package views

import "github.com/sainu/profile-api/pkg/models"

// SocialLinkView is struct of social link view
type SocialLinkView struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewSocialLinksView is constructor for view of social link
func NewSocialLinksView(socialLinks []models.SocialLink) []SocialLinkView {
	var view []SocialLinkView

	for _, socialLink := range socialLinks {
		view = append(view, SocialLinkView{
			Name: socialLink.Name,
			URL:  socialLink.URL,
		})
	}

	return view
}
