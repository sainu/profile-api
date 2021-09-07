package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sainu/profile-api/pkg/models"
	"github.com/sainu/profile-api/pkg/views"
)

// ExperienceController controller for experiences
type ExperienceController struct{}

// NewExperiencesController is constructor for ExperienceController
func NewExperiencesController() *ExperienceController {
	return new(ExperienceController)
}

// Index binds to /experiences
func (c *ExperienceController) Index(ctx echo.Context) error {
	experiences := models.GetAllExperiences()
	view := views.NewExperiencesView(experiences)
	return ctx.JSON(http.StatusOK, view)
}
