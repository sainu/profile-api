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

type queryOption map[string]string

func (o queryOption) apply(req *http.Request) {
	query := req.URL.Query()
	for k, v := range o {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
}

func WithQuery(v map[string]string) queryOption {
	var o queryOption = make(map[string]string)
	for k, v := range v {
		o[k] = v
	}
	return o
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
