package models

import (
	"fmt"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetDetailedFavorites return a list of favorites with it's title, price, etc.
func GetDetailedFavorites(userId int) ([]interfaces.Article, error){
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  // Get favorites list
  var favoriteId string

  query := `SELECT "idArticle" FROM favorites
	    WHERE "idUser" = $1`

  rows, err := conn.QueryContext(ctx, query, userId)
  defer rows.Close()

  if err != nil {
    return []interfaces.Article{}, err
  }

  // Get favorites details
  detailedFavorites := []interfaces.Article{}

  for rows.Next() {
    err = rows.Scan(&favoriteId)
    
    if err != nil {
      return []interfaces.Article{}, err
    }

    mid, err := primitive.ObjectIDFromHex(favoriteId)
    var product interfaces.Article
    err = productsCollection.FindOne(ctx, bson.D{{"_id", mid}}).Decode(&product)

    if err != nil {
      return []interfaces.Article{}, err
    }

    detailedFavorites = append(detailedFavorites, product)
  }

  return detailedFavorites, nil
}

// GetUserOrdersResume Get user orders from database
func GetUserOrdersResume(userId int) ([]interfaces.OrderResume, error){
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  fmt.Println(userId)

  // *** Get orders list
  query := `SELECT "idOrder" FROM orders 
	    WHERE "idUser" = $1`
  
  rows, err := conn.QueryContext(ctx, query, userId)
  defer rows.Close()

  if err != nil {
    return []interfaces.OrderResume{}, err
  }

  // *** Ger orders details
  var orderId int
  orders := []interfaces.OrderResume{}

  for rows.Next() { // For each user order
    err = rows.Scan(&orderId)
    fmt.Println(orderId)

    if err != nil {
      return []interfaces.OrderResume{}, err
    }

    query = `SELECT "idArticle", "amount" from orders_has_products
	    WHERE "idOrder" = $1`
    
    innerRows, err := conn.QueryContext(ctx, query, orderId)
    defer innerRows.Close()

    if err != nil { // Select products on current order
      return []interfaces.OrderResume{}, err
    }

    var products []interfaces.OrderProduct
    var product interfaces.OrderProduct

    for innerRows.Next(){
      err = innerRows.Scan(&product.Product, &product.Amount)

      if err != nil {
	return []interfaces.OrderResume{}, err
      }

      products = append(products, product)
    }

    // Add current order to final array
    orders = append(orders, interfaces.OrderResume{
      Order: orderId, 
      Products: products,
    })
  }

  return orders, nil

}
