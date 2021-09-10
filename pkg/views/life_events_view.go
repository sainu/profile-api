package views

import (
	"time"

	"github.com/sainu/profile-api/pkg/models"
)

type LifeEventView struct {
	Title string     `json:"title"`
	Date  *time.Time `json:"date"`
}

func NewLifeEventsView(lifeEvents *[]models.LifeEvent) *[]LifeEventView {
	var view []LifeEventView

	for _, lifeEvent := range *lifeEvents {
		view = append(view, LifeEventView{
			Title: lifeEvent.Title,
			Date:  lifeEvent.Date,
		})
	}

	return &view
}
