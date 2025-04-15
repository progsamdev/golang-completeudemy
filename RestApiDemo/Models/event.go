package Models

import (
	"errors"
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

func (e *Event) Save() error {

	e.Id = uuid.Must(uuid.NewV4())

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

func GetAllEvents() ([]Event, error) {

	query := `SELECT * FROM events`

	//PREPARE STATEMENT
	preparedStatement, err := db.DBConnection.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer preparedStatement.Close()

	rows, err := preparedStatement.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id uuid.UUID) (*Event, error) {

	query := `SELECT * FROM events WHERE id = ?`

	preparedStatement, err := db.DBConnection.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer preparedStatement.Close()

	var event Event

	row := preparedStatement.QueryRow(id)

	if row == nil {
		return nil, errors.New("event not found")
	}

	err = row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, errors.New("error parsing event")
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `UPDATE events
				SET name = ?, description = ?, location = ?, date_time = ?
				WHERE id = ?`

	preparedStatement, err := db.DBConnection.Prepare(query)
	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	_, err = preparedStatement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	if err != nil {
		return err
	}

	return nil
}

func (event *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	preparedStatement, err := db.DBConnection.Prepare(query)
	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	_, err = preparedStatement.Exec(event.Id)
	if err != nil {
		return err
	}

	return nil
}
