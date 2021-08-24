package views

import "github.com/sainu/profile-api/models"

// WebLinkView is struct of web link view
type WebLinkView struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewWebLinksView is constructor for view of web link
func NewWebLinksView(webLinks []models.WebLink) []WebLinkView {
	var view []WebLinkView

	for _, webLink := range webLinks {
		view = append(view, WebLinkView{
			Name: webLink.Name,
			URL:  webLink.URL,
		})
	}

	return view
}
