package models

import (
	"net/http"
	"time"

	"github.com/sainu/profile-api/pkg/microcms"
)

type LifeEvent struct {
	Title string
	Date  *time.Time
}

func GetAllLifeEvents() *[]LifeEvent {
	req := microcms.NewRequest(
		http.MethodGet,
		"/api/v1/life_events",
		microcms.WithQuery(map[string]string{"orders": "-date"}),
	)
	resp := new(microcms.LifeEventsResponseBody)
	client := microcms.NewClient()
	client.Do(req, resp)

	var lifeEvents []LifeEvent
	for _, c := range resp.Contents {
		date, _ := time.Parse(time.RFC3339, c.Date)
		lifeEvents = append(lifeEvents, LifeEvent{
			Title: c.Title,
			Date:  &date,
		})
	}

	return &lifeEvents
}
