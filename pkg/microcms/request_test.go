package microcms

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithQuery(t *testing.T) {
	actual := WithQuery(map[string]string{"key1": "val1", "key2": "val2"})
	expected := queryOption{"key1": "val1", "key2": "val2"}
	assert.Equal(t, expected, actual)
}

func TestNewRequest(t *testing.T) {
	t.Run("it set X-API-KEY header", func(t *testing.T) {
		req := NewRequest(http.MethodGet, "/path")
		assert.Equal(t, apiKey, req.Header.Get("X-API-KEY"))
	})

	t.Run("it set query", func(t *testing.T) {
		req := NewRequest(
			http.MethodGet,
			"/path",
			WithQuery(map[string]string{"key1": "val1", "key2": "val2"}),
		)
		assert.Equal(t, "key1=val1&key2=val2", req.URL.RawQuery)
	})
}
