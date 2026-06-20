package database

import (
	"context"
	"database/sql"
	"time"
)

type EventModels struct {
	DB *sql.DB
}

type Event struct {
	Id          int       `json:"id"`
	OwnerId     int       `json:"OwnerId" binding:"required"`
	Name        string    `json:"name" binding:"required, min=3"`
	Description string    `json:"description"  binding:"required , min=3 , max=10"`
	Date        time.Time `json:"date" binding:"required , datetime=2006-01-02"`
	Location    string    `json:"location"  binding:"required , min=3"`
}

func (em *EventModels) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (OwnerId , name , description , data , location) VALUE($1,$2,$3,$4,$5)"
	return em.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(&event.Id)

}
func (em *EventModels) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	query := "SELECT * FROM events"
	rows, err := em.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []*Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(event.Id, event.OwnerId, event.Name, event.Description, event.Date, event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
func (em *EventModels) Get(id int) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events WHERE id = $1"
	var event Event
	err := em.DB.QueryRowContext(ctx, query, id).Scan(event.Id, event.OwnerId, event.Name, event.Description, event.Date, event.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}
func (em *EventModels) Update(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE events name = $1 , description = $2 , date = $3 , location = $4 WHERE id = $5"
	_, err := em.DB.ExecContext(ctx, query, event.Name, event.Description, event.Date, event.Location, event.Id)
	if err != nil {
		return err
	}
	return nil
}
func (em *EventModels) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM events WHERE id = $1"

	_, err := em.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
