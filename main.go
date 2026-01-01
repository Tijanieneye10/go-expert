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

func main() {
	db, err := sql.Open("sqlite3", "database.sql")

	if err != nil {
		log.Fatal(err)
	}

	myUser, err := GetUserByEmail("john@gmail.com", db)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(myUser.Username)

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	user := User{
		Id:       1,
		Username: "john",
		Email:    "john2@gmail.com",
		Password: string(passwordHashed),
	}

	err = AddUser(&user, db)

	if err != nil {
		log.Fatal(err)
	}

	//createTable(db)

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

func AddUser(user *User, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO users(username, email, password) values(?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password)

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
