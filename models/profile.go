package models

import "fmt"

// Profile is struct of profile
type Profile struct {
	FamilyNameKanji string `json:"family_name_kanji"`
	GivenNameKanji  string `json:"given_name_kanji"`
	FamilyNameKana  string `json:"family_name_kana"`
	GivenNameKana   string `json:"given_name_kana"`
	FamilyNameEn    string `json:"family_name_en"`
	GivenNameEn     string `json:"given_name_en"`
	Nickname        string `json:"nickname"`
	Job             string `json:"job"`
	Email           string `json:"email"`
	Bio             string `json:"bio"`
	Location        string `json:"location"`
}

// GetProfile returns a profile
func GetProfile() *Profile {
	return &Profile{
		FamilyNameKanji: "道祖",
		GivenNameKanji:  "克理",
		FamilyNameKana:  "さいのう",
		GivenNameKana:   "かつとし",
		FamilyNameEn:    "Saino",
		GivenNameEn:     "Katsutoshi",
		Nickname:        "sainu",
		Job:             "Software Program Developer",
		Email:           "katsutoshi.saino@gmail.com",
		Bio:             "東京都出身のエンジニア。大学在学中にインターンや受託でWeb開発を経験。その後、個人開発を経て、お金の会社に就職。Railsでのサーバー開発から始まり、JSのFW(AngularやNuxt)を使ったフロントエンド開発、AWSでインフラの構築を経験してきました。",
		Location:        "Tokyo",
	}
}

// FullNameKanji is full name in Japanese
func (p *Profile) FullNameKanji() string {
	return fmt.Sprintf("%s %s", p.FamilyNameKanji, p.GivenNameKanji)
}

// FullNameKana is kana in Japanese
func (p *Profile) FullNameKana() string {
	return fmt.Sprintf("%s %s", p.FamilyNameKana, p.GivenNameKana)
}

// FullNameEn is full name in English
func (p *Profile) FullNameEn() string {
	return fmt.Sprintf("%s %s", p.GivenNameEn, p.FamilyNameEn)
}
