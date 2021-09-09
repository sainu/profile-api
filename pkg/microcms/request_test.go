package microcms

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	t.Run("it set X-API-KEY header", func(t *testing.T) {
		req := NewRequest(http.MethodGet, "/path")
		assert.Equal(t, apiKey, req.Header.Get("X-API-KEY"))
	})
}
