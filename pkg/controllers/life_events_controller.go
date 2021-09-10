package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sainu/profile-api/pkg/models"
	"github.com/sainu/profile-api/pkg/views"
)

type LifeEventsController struct{}

func NewLifeEventsController() *LifeEventsController {
	return new(LifeEventsController)
}

func (c *LifeEventsController) Index(ctx echo.Context) error {
	lifeEvents := models.GetAllLifeEvents()
	view := views.NewLifeEventsView(lifeEvents)
	return ctx.JSON(http.StatusOK, view)
}
