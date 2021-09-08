package microcms

import (
	"net/http"
	"os"
)

var (
	baseURL = os.Getenv("MICRO_CMS_PROFILE_API_BASE_URL")
	apiKey  = os.Getenv("MICRO_CMS_X_API_KEY")
)

func NewRequest(path string) *http.Request {
	req, _ := http.NewRequest(
		http.MethodGet,
		baseURL+path,
		nil,
	)
	req.Header.Set("X-API-KEY", apiKey)
	return req
}