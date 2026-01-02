package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int       `json:"id"`
	Username string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password" db:"_"`
	Created  time.Time `json:"created"`
}

type Profile struct {
	Id      int     `json:"id"`
	UserId  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}

func main() {
	db, err := sql.Open("sqlite3", "database.sql")

	err = createTable(db)
	if err != nil {
		return
	}

	//myUser, err := GetUserByEmail("john@gmail.com", db)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(myUser.Username)

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	user := User{
		Id:       1,
		Username: "john",
		Email:    "john10@gmail.com",
		Password: string(passwordHashed),
	}

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	userId, err := AddUser(&user, tx)

	err = AddProfile(userId, tx)
	if err != nil {
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

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

func AddUser(user *User, db *sql.Tx) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users(username, email, password) values(?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	data, err := stmt.Exec(user.Username, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	return data.LastInsertId()
}

func AddProfile(userId int64, tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO profiles(user_id, balance) values(?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, 100.00)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string, db *sql.DB) (*User, error) {
	query, err := db.Prepare("SELECT * FROM users WHERE email=?")

	if err != nil {
		return nil, err
	}

	defer query.Close()

	row := query.QueryRow(email)

	var user User

	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Created)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func createTable(db *sql.DB) error {
	// 1. Create Users Table
	// note: In SQLite, 'INTEGER PRIMARY KEY' automatically auto-increments.
	queryUsers := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY, 
        username TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

	if _, err := db.Exec(queryUsers); err != nil {
		return err
	}

	// 2. Create Profiles Table
	queryProfiles := `
    CREATE TABLE IF NOT EXISTS profiles (
        id INTEGER PRIMARY KEY,
        user_id INTEGER NOT NULL,
        balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`

	if _, err := db.Exec(queryProfiles); err != nil {
		return err
	}

	return nil
}
