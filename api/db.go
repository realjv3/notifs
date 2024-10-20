package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const createUsers string = `
CREATE TABLE users (
	id TEXT NOT NULL PRIMARY KEY,
	email TEXT NOT NULL,
	preference TEXT NOT NULL
);`
const createNotifications string = `
CREATE TABLE notifications (
	id TEXT NOT NULL PRIMARY KEY,
	userId INTEGER NOT NULL,
	title TEXT NOT NULL,
	description TEXT,
	type TEXT NOT NULL,
	createdAt DATETIME NOT NULL,
	FOREIGN KEY (userId)
       REFERENCES users (userId)
);`

var db *sqlx.DB

// Initialize an in-memory SQLite database, create tables, and seed data
func initDB() *sqlx.DB {
	var err error

	fmt.Print("Connecting to the database... ")
	db, err = sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		fmt.Print("Failed to connect to DB", err)
		os.Exit(1)
	}
	fmt.Println("Done")

	fmt.Print("Creating users table... ")
	if _, err = db.Exec(createUsers); err != nil {
		fmt.Print("Failed to create users table", err)
		os.Exit(1)
	}
	fmt.Println("Done")

	fmt.Print("Creating notifications table... ")
	if _, err = db.Exec(createNotifications); err != nil {
		fmt.Print("Failed to create notifications table", err)
		os.Exit(1)
	}
	fmt.Println("Done")

	fmt.Print("Seeding the users table... ")
	execSQLFile(db, "./seed/users.sql")
	fmt.Println("Done")

	fmt.Print("Seeding the notifications table... ")
	execSQLFile(db, "./seed/notifications.sql")
	fmt.Println("Done")

	return db
}

func execSQLFile(db *sqlx.DB, sqlFilePath string) {
	seedNotifications, err := os.ReadFile(sqlFilePath)
	if err != nil {
		fmt.Print("Failed to load file", sqlFilePath, err)
		os.Exit(1)
	}

	if _, err := db.Exec(string(seedNotifications)); err != nil {
		fmt.Print("Failed to execute file", sqlFilePath, err)
		os.Exit(1)
	}
}
