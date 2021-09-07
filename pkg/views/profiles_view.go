package views

import "github.com/sainu/profile-api/pkg/models"

// ProfileView is struct of profile view
type ProfileView struct {
	FamilyNameKanji string `json:"family_name_kanji"`
	GivenNameKanji  string `json:"given_name_kanji"`
	FamilyNameKana  string `json:"family_name_kana"`
	GivenNameKana   string `json:"given_name_kana"`
	FamilyNameEn    string `json:"family_name_en"`
	GivenNameEn     string `json:"given_name_en"`
	FullNameKanji   string `json:"full_name_kanji"`
	FullNameKana    string `json:"full_name_kana"`
	FullNameEn      string `json:"full_name_en"`
	Nickname        string `json:"nickname"`
	Job             string `json:"job"`
	Email           string `json:"email"`
	Bio             string `json:"bio"`
	Location        string `json:"location"`
}

// NewProfileView is constructor for view of profile
func NewProfileView(profile *models.Profile) *ProfileView {
	view := &ProfileView{
		FamilyNameKanji: profile.FamilyNameKana,
		GivenNameKanji:  profile.GivenNameKanji,
		FamilyNameKana:  profile.FamilyNameKana,
		GivenNameKana:   profile.GivenNameKana,
		FamilyNameEn:    profile.FamilyNameEn,
		GivenNameEn:     profile.GivenNameEn,
		FullNameKanji:   profile.FullNameKanji(),
		FullNameKana:    profile.FullNameKana(),
		FullNameEn:      profile.FullNameEn(),
		Nickname:        profile.Nickname,
		Job:             profile.Job,
		Email:           profile.Email,
		Bio:             profile.Bio,
		Location:        profile.Location,
	}
	return view
}
