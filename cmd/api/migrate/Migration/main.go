package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Args most be up or down")

	}
	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("can not connect to database")

	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fSrc, err := (&file.File{}).Open("file://cmd/migrate/migration")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithInstance(
		"file",
		fSrc,
		"sqlite3",
		instance,
	)
	if err != nil {
		log.Fatal(err)
	}
	switch direction {
	case "Up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "Down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)

		}
	default:
		log.Fatal("invalid diraction use 'up' or  'down'")

	}
}
