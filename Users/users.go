package users

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var dbUrl = "dbusername:password@tcp(ip:port)/dbname"

func CreateNewUser(username string, password string) bool {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println("Error in connection")
		fmt.Println(err)
	}
	defer db.Close()

	var fetchedUsername string
	db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&fetchedUsername)

	if fetchedUsername != username {
		if username != "" && password != "" {
			db.QueryRow("INSERT INTO users (username, password) VALUES (?,?)", username, hashPassword(password))
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func GetUser(username string, password string) bool {
	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		fmt.Println("Error in connection")
		fmt.Println(err)
	}
	defer db.Close()

	var fetchedUsername string
	getUsername := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&fetchedUsername)

	if getUsername == sql.ErrNoRows {
		return false
	} else if fetchedUsername == username {
		if checkPassword(fetchedUsername, password) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func checkPassword(username string, password string) bool {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println("Error in connection")
		fmt.Println(err)
	}
	defer db.Close()

	var fetchedPassword string
	getPassword := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&fetchedPassword)

	if bcrypt.CompareHashAndPassword([]byte(fetchedPassword), []byte(password)) == nil && getPassword != sql.ErrNoRows {
		return true
	} else {
		return false
	}
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("Could not hash password")
		fmt.Println(err)
	}
	return string(bytes)
}
