package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create user table")
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
		FOREIGN KEY (userId) REFERENCES user(id)
	)
	`

	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create event table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registration (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	eventId INTEGER,
	userId INTEGER,
	FOREIGN KEY (eventId) REFERENCES event(id),
	FOREIGN KEY (userId) REFERENCES user(id)
	)
	`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("Could not create registration table")
	}
}
