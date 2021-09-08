package models

import (
	"github.com/sainu/profile-api/pkg/microcms"
)

// SocialLink is struct of social link
type SocialLink struct {
	Name string
	URL  string
}

// GetAllSocialLinks returns all social links
func GetAllSocialLinks() *[]SocialLink {
	resp := microcms.GetSocialLinks()
	var socialLinks []SocialLink
	for _, c := range resp.Contents {
		socialLinks = append(socialLinks, SocialLink{
			Name: c.Name,
			URL:  c.Url,
		})
	}
	return &socialLinks
}
