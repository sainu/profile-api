package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sainu/profile-api/pkg/models"
	"github.com/sainu/profile-api/pkg/views"
)

// SkillsController controller for skill
type SkillsController struct{}

// NewSkillsController is constructor for SkillsController
func NewSkillsController() *SkillsController {
	return new(SkillsController)
}

// Index binds to /skills
func (c *SkillsController) Index(ctx echo.Context) error {
	skills := models.GetAllSkills()
	view := views.NewSkillsView(*skills)
	return ctx.JSON(http.StatusOK, view)
}
