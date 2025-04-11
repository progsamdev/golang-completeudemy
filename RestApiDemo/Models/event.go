package Models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      uuid.UUID `json:"user_id"`
}

var events []Event

func (e *Event) Save() error {
	events = append(events, *e)
	return nil
}

func GetAllEvents() []Event {
	return events
}
