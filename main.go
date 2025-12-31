package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sql")

	if err != nil {
		log.Fatal(err)
	}

	createTable(db)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		return
	}

	fmt.Println("Connected to database")
}

func createTable(db *sql.DB) {
	smtm := `
	CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
	password VARCHAR(250),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

	_, err := db.Exec(smtm)

	if err != nil {
		log.Fatal(err)
	}
}
