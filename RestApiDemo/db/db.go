package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DBConnection *sql.DB

func InitDB() (*sql.DB, error) {
	fmt.Println("Opening database connection...")
	var err error
	DBConnection, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		fmt.Println("Error opening database:", err)
		panic("Failed to connect to database")
	}

	if err := DBConnection.Ping(); err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}
	DBConnection.SetMaxOpenConns(10)
	DBConnection.SetMaxIdleConns(5)

	fmt.Println("Creating table...")
	createTable()

	return DBConnection, nil
}

func createTable() {

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY, 
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DBConnection.Exec(createUserTable)
	if err != nil {
		panic("Failed to create table: " + err.Error())
	}

	sqlCreateTable := `
	CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		location TEXT,
		date_time DATETIME NOT NULL,
		user_id TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	_, err = DBConnection.Exec(sqlCreateTable)
	if err != nil {
		panic("Failed to create table: " + err.Error())
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id TEXT PRIMARY KEY,
		event_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	)
	`

	_, err = DBConnection.Exec(createRegistrationTable)
	if err != nil {
		panic("Failed to create table: " + err.Error())
	}
}
