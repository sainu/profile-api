package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sainu/profile-api/models"
	"github.com/sainu/profile-api/views"
)

// Handler is handler function for vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	profile := models.GetProfile()
	view := views.NewProfileView(profile)
	bytes, _ := json.Marshal(view)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(bytes))
}
