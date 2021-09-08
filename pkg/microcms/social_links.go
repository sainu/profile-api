package microcms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func GetSocialLinks() *SocialLinksResponseBody {
	req := NewRequest("/api/v1/social_links")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	unmarshalResp := new(SocialLinksResponseBody)
	if err := json.Unmarshal(bytes, unmarshalResp); err != nil {
		panic(err)
	}
	return unmarshalResp
}
