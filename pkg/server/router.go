package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sainu/profile-api/pkg/controllers"
)

// NewRouter is constructor for router
func NewRouter() (*echo.Echo, error) {
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.BasicAuth(BasicAuthHandler))

	experiencesController := controllers.NewExperiencesController()
	router.GET("/api/experiences", experiencesController.Index)

	lifeEventsController := controllers.NewLifeEventsController()
	router.GET("/api/life_events", lifeEventsController.Index)

	profilesController := controllers.NewProfilesController()
	router.GET("/api/profile", profilesController.Show)

	skillsController := controllers.NewSkillsController()
	router.GET("/api/skills", skillsController.Index)

	socialLinksController := controllers.NewSocialLinksController()
	router.GET("/api/social_links", socialLinksController.Index)

	return router, nil
}
