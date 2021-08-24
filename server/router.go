package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sainu/profile-api/controllers"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
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
