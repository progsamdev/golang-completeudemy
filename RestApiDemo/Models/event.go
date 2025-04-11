package Models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      uuid.UUID `json:"user_id"`
}
