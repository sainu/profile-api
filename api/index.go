package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sainu/profile-api/controllers"
)

// Handler is handler function for vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router, err := newRouter()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	router.ServeHTTP(w, r)
}

// NewRouter is constructor for router
func newRouter() (*echo.Echo, error) {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	profilesController := controllers.NewProfilesController()
	router.GET("/api/profile", profilesController.Show)

	skillsController := controllers.NewSkillsController()
	router.GET("/api/skills", skillsController.Index)

	socialLinksController := controllers.NewSocialLinksController()
	router.GET("/api/social_links", socialLinksController.Index)

	webLinksController := controllers.NewWebLinksController()
	router.GET("/api/web_links", webLinksController.Index)

	return router, nil
}
