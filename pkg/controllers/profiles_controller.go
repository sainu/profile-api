package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sainu/profile-api/pkg/models"
	"github.com/sainu/profile-api/pkg/views"
)

// ProfilesController controller for profile
type ProfilesController struct{}

// NewProfilesController is constructor for ProfileController
func NewProfilesController() *ProfilesController {
	return new(ProfilesController)
}

// Show binds to /profile
func (c *ProfilesController) Show(ctx echo.Context) error {
	profile := models.GetProfile()
	view := views.NewProfileView(profile)
	return ctx.JSON(http.StatusOK, view)
}
