package models

import (
	"context"
	"log"
	"time"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/SilviaPabon/buenavida-backend/interfaces"
)

// Connection db
var conn = configs.ConnectToPostgres()

// SaveUser Create an user on database
func SaveUser(u *interfaces.Users) (r bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO users (name, lastname, mail, password)
              VALUES ($1, $2, $3, $4);
            `
	row := conn.QueryRowContext(
		ctx, query, u.Firstname, u.Lastname, u.Email, u.Password,
	)

	if row.Err() != nil {
		log.Fatal(row.Err())
	}

	return true
}

// FindByEmail search an user by mail
func FindByEmail(email string) (succ bool) {

	// Search on database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(*) AS count FROM users
    		WHERE UPPER(users.mail) = UPPER($1);`

	row := conn.QueryRowContext(
		ctx, query, email,
	)
	//this checks if a email already exists
	var check int
	err := row.Scan(&check)
	if err != nil {
		log.Fatal(err)
	}

	if check == 1 {
		return true
	}

	return false
}
