package models

import (
	"fmt"

	"github.com/sainu/profile-api/pkg/microcms"
)

// Profile is struct of profile
type Profile struct {
	FamilyNameKanji string
	GivenNameKanji  string
	FamilyNameKana  string
	GivenNameKana   string
	FamilyNameEn    string
	GivenNameEn     string
	Nickname        string
	Job             string
	Email           string
	Bio             string
	Location        string
}

// GetProfile returns a profile
func GetProfile() *Profile {
	client := microcms.NewClient()
	resp := new(microcms.ProfileResponseBody)
	client.Do(microcms.NewRequest("/api/v1/profile"), resp)

	profile := &Profile{
		FamilyNameKanji: resp.FamilyNameKanji,
		GivenNameKanji:  resp.GivenNameKanji,
		FamilyNameKana:  resp.FamilyNameKana,
		GivenNameKana:   resp.GivenNameKana,
		FamilyNameEn:    resp.FamilyNameEn,
		GivenNameEn:     resp.GivenNameEn,
		Nickname:        resp.Nickname,
		Job:             resp.Job,
		Email:           resp.Email,
		Bio:             resp.Bio,
		Location:        resp.Location,
	}

	return profile
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
