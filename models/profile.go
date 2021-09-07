package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sainu/profile-api/pkg/microcms"
)

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
	req := microcms.NewRequest("/api/v1/profile")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	profile := new(Profile)
	if err := json.Unmarshal(byteArray, profile); err != nil {
		panic(err)
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
