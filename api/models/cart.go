package models

import(
  // "fmt"
  "context"
  "time"
  "github.com/SilviaPabon/buenavida-backend/configs"
  // "github.com/SilviaPabon/buenavida-backend/interfaces"
)

// Mongodb collections
var pg = configs.ConnectToPostgres()

// AddProductToCart Add a new product to the cart or update the current ammount
func AddProductToCart(userId int, productId string) error {
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  // ### ### ###
  // Verify if the product already exists
  query := `SELECT COUNT("idUser") AS count FROM cart WHERE
	    "idUser" = $1 AND "idArticle" = $2;`
  
  row := pg.QueryRowContext(ctx, query, userId, productId)
  var count int
  err := row.Scan(&count)

  if err != nil{
    return err
  }

  // ### #### ### Insert or update as needed
  if count == 0 {
    query = `INSERT INTO cart ("idUser", "idArticle", "amount")
	      VALUES ($1, $2, $3);` 
    row = pg.QueryRowContext(ctx, query, userId, productId, 1)
    return row.Err() // Returns nil or error if some error exists
  }else{
    query = `UPDATE cart 
	    SET "amount" = "amount" + 1
	    WHERE "idUser" = $1 AND "idArticle" = $2`
    row = pg.QueryRowContext(ctx, query, userId, productId)
    return row.Err()
  }

}

// UpdateProductInCart Update the amount of some product on database
func UpdateProductInCart(userId, amount int, productId string) error {
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()
  
  // ### ### ###
  // Verify if the product already exists
  query := `SELECT COUNT("idUser") AS count FROM cart WHERE
	    "idUser" = $1 AND "idArticle" = $2;`
  
  row := pg.QueryRowContext(ctx, query, userId, productId)
  var count int
  err := row.Scan(&count)

  if err != nil{
    return err
  }

  // ### ### ### Insert or update as needed
  if count == 0 {
    query = `INSERT INTO cart ("idUser", "idArticle", "amount")
	    VALUES ($1, $2, $3);`
    row = pg.QueryRowContext(ctx, query, userId, productId, amount)
    return row.Err()
  }else{
    query = `UPDATE cart
	    SET "amount" = $1
	    WHERE "idUser" = $2 AND "idArticle" = $3`
    row = pg.QueryRowContext(ctx, query, amount, userId, productId)
    return row.Err()
  }
}

// GetCartLength Gets the user cart lenght
func GetCartLength(userId int) (int, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  // Prepare query
  query := `SELECT COUNT("idUser") AS count FROM CART WHERE
	    "idUser" = $1;`
  
  // Make the request
  row := pg.QueryRowContext(ctx, query, userId)
  var count int
  err := row.Scan(&count)

  if err != nil {
    return 0, err
  }

  return count, nil
}

// CreateOrder Calls the stored procedure and creates the order from the user cart
func CreateOrder(userId int) (error) {
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  // Prepare query
  query := `CALL make_order($1)`
  row := pg.QueryRowContext(ctx, query, userId)
  return row.Err() // Returns error if any
}
