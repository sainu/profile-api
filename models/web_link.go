package models

// WebLink is struct of web link
type WebLink struct {
	Name string
	URL  string
}

// GetAllWebLinks returns all web links
func GetAllWebLinks() *[]WebLink {
	return &[]WebLink{
		{
			Name: "職務経歴",
			URL:  "https://github.com/sainu/resume",
		},
		{
			Name: "ブログ",
			URL:  "https://sainu.hatenablog.jp/",
		},
	}
}
