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
	client := microcms.NewClient()
	resp := new(microcms.SocialLinksResponseBody)
	client.Do(microcms.NewRequest("/api/v1/social_links"), resp)
	var socialLinks []SocialLink
	for _, c := range resp.Contents {
		socialLinks = append(socialLinks, SocialLink{
			Name: c.Name,
			URL:  c.Url,
		})
	}
	return &socialLinks
}
