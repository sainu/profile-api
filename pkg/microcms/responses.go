package microcms

type ProfileResponseBody struct {
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	PublishedAt     string `json:"publishedAt"`
	RevisedAt       string `json:"revisedAt"`
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

type SocialLinksResponseBody struct {
	Contents [](struct {
		Id          string `json:"id"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		PublishedAt string `json:"publishedAt"`
		RevisedAt   string `json:"revisedAt"`
		Name        string `json:"name"`
		Url         string `json:"url"`
	}) `json:"contents"`
	TotalCount int8 `json:"totalCount"`
	Offset     int8 `json:"offset"`
	Limit      int8 `json:"limit"`
}

type SkillsResponseBody struct {
	Contents [](struct {
		Id          string `json:"id"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		PublishedAt string `json:"publishedAt"`
		RevisedAt   string `json:"revisedAt"`
		Name        string `json:"name"`
		Score       uint8  `json:"score"`
	}) `json:"contents"`
	TotalCount int8 `json:"totalCount"`
	Offset     int8 `json:"offset"`
	Limit      int8 `json:"limit"`
}

type ExperienceResponseBody struct {
	Contents [](struct {
		Id              string   `json:"id"`
		CreatedAt       string   `json:"createdAt"`
		UpdatedAt       string   `json:"updatedAt"`
		PublishedAt     string   `json:"publishedAt"`
		RevisedAt       string   `json:"revisedAt"`
		CompanyName     string   `json:"companyName"`
		EmploymentTypes []string `json:"employmentTypes"`
		StartDate       string   `json:"startDate"`
		EndDate         string   `json:"endDate"`
		Projects        [](struct {
			FieldId      string `json:"fieldId"`
			Description  string `json:"description"`
			Technologies [](struct {
				FieldId    string `json:"fieldId"`
				Technology struct {
					Id          string `json:"id"`
					CreatedAt   string `json:"createdAt"`
					UpdatedAt   string `json:"updatedAt"`
					PublishedAt string `json:"publishedAt"`
					RevisedAt   string `json:"revisedAt"`
					Name        string `json:"name"`
				} `json:"technology"`
				Versions string `json:"versions"`
			}) `json:"technologies"`
		}) `json:"projects"`
	}) `json:"contents"`
	TotalCount int8 `json:"totalCount"`
	Offset     int8 `json:"offset"`
	Limit      int8 `json:"limit"`
}
