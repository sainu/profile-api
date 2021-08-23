package models

// WebLink is struct of web link
type WebLink struct {
	Title string
	URL   string
}

// GetAllWebLinks returns all web links
func GetAllWebLinks() *[]WebLink {
	return &[]WebLink{
		{
			Title: "職務経歴",
			URL:   "https://github.com/sainu/resume",
		},
		{
			Title: "ブログ",
			URL:   "https://sainu.hatenablog.jp/",
		},
	}
}
