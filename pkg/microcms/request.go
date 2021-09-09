package microcms

import (
	"net/http"
	"os"
)

var (
	baseURL = os.Getenv("MICRO_CMS_PROFILE_API_BASE_URL")
	apiKey  = os.Getenv("MICRO_CMS_X_API_KEY")
)

type option interface {
	apply(*http.Request)
}

func NewRequest(method string, path string, options ...option) *http.Request {
	req, _ := http.NewRequest(
		method,
		baseURL+path,
		nil,
	)
	req.Header.Set("X-API-KEY", apiKey)
	for _, o := range options {
		o.apply(req)
	}
	return req
}
