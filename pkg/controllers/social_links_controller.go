package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sainu/profile-api/pkg/models"
	"github.com/sainu/profile-api/pkg/views"
)

// SocialLinksController controller for social link
type SocialLinksController struct{}

// NewSocialLinksController is constructor for SocialLinksController
func NewSocialLinksController() *SocialLinksController {
	return new(SocialLinksController)
}

// Index binds to /social_links
func (c *SocialLinksController) Index(ctx echo.Context) error {
	socialLinks := models.GetAllSocialLinks()
	view := views.NewSocialLinksView(*socialLinks)
	return ctx.JSON(http.StatusOK, view)
}
