package database

import "database/sql"

type Models struct {
	User      UserModels
	Event     EventModels
	Attendees attendeesModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		User:      UserModels{DB: db},
		Event:     EventModels{DB: db},
		Attendees: attendeesModel{DB: db},
	}
}
