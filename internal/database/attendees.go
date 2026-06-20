package database

import "database/sql"

type attendeesModel struct {
	DB *sql.DB
}

type Attendees struct {
	Id      int `json:"id"`
	UserId  int `json:"Userid"`
	EventId int `json:"EventID"`
}
