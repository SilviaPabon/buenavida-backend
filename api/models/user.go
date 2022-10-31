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

	/* _, err := usersCollection.InsertOne(ctx, u)

	   if err != nil{
	     return false
	   } */

	query := `INSERT INTO users (name, lastname, mail, password)
              VALUES ($1, $2, $3, $4);
            `
	row := conn.QueryRowContext(
		ctx, query, u.Firstname, u.Lastname, u.Email, u.Password,
	)
	log.Println(row)
	return true
}
