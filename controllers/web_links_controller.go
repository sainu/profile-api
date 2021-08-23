package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sainu/profile-api/models"
	"github.com/sainu/profile-api/views"
)

// WebLinksController controller for web link
type WebLinksController struct{}

// NewWebLinksController is constructor for WebLinksController
func NewWebLinksController() *WebLinksController {
	return new(WebLinksController)
}

// Index binds to /web_links
func (c *WebLinksController) Index(ctx echo.Context) error {
	webLinks := models.GetAllWebLinks()
	view := views.NewWebLinksView(*webLinks)
	return ctx.JSON(http.StatusOK, view)
}
