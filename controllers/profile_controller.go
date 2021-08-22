package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sainu/profile-api/models"
	"github.com/sainu/profile-api/views"
)

// ProfileController controller for profile
type ProfileController struct{}

// NewProfileController is constructor for ProfileController
func NewProfileController() *ProfileController {
	return new(ProfileController)
}

// Show binds to /profile
func (c *ProfileController) Show(ctx echo.Context) error {
	profile := models.GetProfile()
	view := views.NewProfileView(profile)
	return ctx.JSON(http.StatusOK, view)
}
