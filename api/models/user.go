package models

import (
	// "fmt"
	"context"
	"errors"
	"log"
	"time"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/SilviaPabon/buenavida-backend/interfaces"
)

// Connection db
var conn = configs.ConnectToPostgres()

// SaveUser Create an user on database
func SaveUser(u *interfaces.User) (r bool) {

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

// GetUserFromMail get firstname, lastname, mail and password from given mail
func GetUserFromMail(mail string) (interfaces.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user interfaces.User

	// Get user from database
	query := `SELECT id, name, lastname, mail, password FROM USERS WHERE UPPER(USERS.mail) = UPPER($1) LIMIT 1`

	rows, err := conn.QueryContext(ctx, query, mail)
	defer rows.Close()

	if err != nil {
		return interfaces.User{}, err
	}

	// Parse user
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password)

		if err != nil {
			return interfaces.User{}, err
		}
	}

	if user.ID == 0 || user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Password == "" {
		return interfaces.User{}, errors.New("Not found")
	}

	return user, nil
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

func AddFavorites(userId int, idArticle string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verify if the product already exists
	query := `SELECT COUNT("idUser") AS count FROM favorites WHERE
	    "idUser" = $1 AND "idArticle" = $2;`

	row := conn.QueryRowContext(ctx, query, userId, idArticle)
	var count int
	err := row.Scan(&count)

	if err != nil {
		return err
	}

	// ### #### ### Insert or update as needed
	if count == 0 {
		query = `INSERT INTO favorites ("idUser", "idArticle")
	    VALUES ($1, $2);`
		row = conn.QueryRowContext(ctx, query, userId, idArticle)
		return row.Err() // Returns nil or error if some error exists
	} else {
		return interfaces.ErrAlreadyInFavorites
	}
}

func FavoritesGET(idUser int) ([]string, error) {
	//Search DataBase
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Favoritesgetid interfaces.Favorite

	query := `SELECT "idArticle" FROM favorites WHERE "idUser" = $1`

	row, err := conn.QueryContext(ctx, query, idUser)
	defer row.Close()

	if err != nil {
		return []string{}, err
	}

	favoriteArray := []string{}
	for row.Next() {
		err = row.Scan(&Favoritesgetid.FavoriteId)
		favoriteArray = append(favoriteArray, Favoritesgetid.FavoriteId)
		if err != nil {
			return []string{}, err
		}
	}

	return favoriteArray, nil
}
