package Models

import (
	"fmt"
	"restapidemo/db"
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      uuid.UUID `json:"user_id"`
}

var events []Event

func (e *Event) Save() error {

	e.Id = uuid.Must(uuid.NewV4())
	e.UserID = uuid.Must(uuid.NewV4())

	query := `INSERT INTO events (id, name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?, ?)`

	preparedStatement, err := db.DBConnection.Prepare(query)
	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	result, err := preparedStatement.Exec(e.Id, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(lastInsertId)

	return nil
}

func GetAllEvents() []Event {
	return events
}
