package views

import "github.com/sainu/profile-api/models"

// ProfileView is struct of profile view
type ProfileView struct {
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

// NewProfileView is constructor for view of profile
func NewProfileView(profile *models.Profile) *ProfileView {
	view := &ProfileView{
		FamilyNameKanji: profile.FamilyNameKana,
		GivenNameKanji:  profile.GivenNameKanji,
		FamilyNameKana:  profile.FamilyNameKana,
		GivenNameKana:   profile.GivenNameKana,
		FamilyNameEn:    profile.FamilyNameEn,
		GivenNameEn:     profile.GivenNameEn,
		Nickname:        profile.Nickname,
		Job:             profile.Job,
		Email:           profile.Email,
		Bio:             profile.Bio,
		Location:        profile.Location,
	}
	return view
}
