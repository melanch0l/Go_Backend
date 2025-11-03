package models

import (
	"example/restapi/db"
	"fmt"
	"time"
)

type Event struct {
	ID int64
	//using struct tag to enfore required data
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name,description,location,datetime,user_id)
	VALUES (?,?,?,?,?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	//use Exec for writing
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}
func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?,description = ?,location = ?, datetime = ?
	WHERE id =?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.Name, e.Description, e.Location, e.Datetime, e.ID)
	return err
}

// delete event from db
func (e *Event) Delete() error {
	query := `
	DELETE FROM events WHERE id = ?
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID)
	return err
}
func (e *Event) Register(userID int64) error {
	fmt.Println(userID)
	query := `
	INSERT INTO registrations(event_id,user_id)
	VALUES(?,?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID, userID)

	return err
}
func (e *Event) Cancel(userID int64) error {
	query := `
	DELETE FROM registrations WHERE
	event_id = ? AND user_id = ? 
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID, userID)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events`
	//use Query for viewing
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}
func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events 
	WHERE "id" = ?
	`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
