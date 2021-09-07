package handler

import (
	"net/http"

	"github.com/sainu/profile-api/pkg/server"
)

// Handler is handler function for vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router, err := server.NewRouter()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	router.ServeHTTP(w, r)
}
