package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/pelletier/go-toml/query"
)

type attendeesModel struct {
	DB *sql.DB
}

type Attendees struct {
	Id      int `json:"id"`
	UserId  int `json:"Userid"`
	EventId int `json:"EventID"`
}

func (m *attendeesModel) Instert(attendees *Attendees) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO attendees (event_id , user_id) VALUES($1 , $2) RETURNING id"
	


}
