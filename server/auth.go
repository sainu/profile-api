package server

import (
	"crypto/subtle"
	"os"

	"github.com/labstack/echo/v4"
)

// BasicAuthHandler is handler for basic auth
func BasicAuthHandler(username, password string, c echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte(os.Getenv("BASIC_AUTH_USERNAME"))) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(os.Getenv("BASIC_AUTH_PASSWORD"))) == 1 {
		return true, nil
	}
	return false, nil
}
