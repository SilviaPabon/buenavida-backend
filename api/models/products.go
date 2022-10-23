package models

import(
  // "fmt"
  "context"
  "time"
  "go.mongodb.org/mongo-driver/bson"
  //"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/SilviaPabon/buenavida-backend/configs"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

// Mongodb collection
var productsCollection = configs.GetCollection("products")

// GetProductsByPage get products by given page
func GetProductsByPage(page int) (p []interfaces.Article, e error) {
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  // Prepare query
  options := options.Find().SetSkip(int64(page -1) * 5).SetLimit(5)

  // Make query
  var products []interfaces.Article
  cursor, _ := productsCollection.Find(ctx, bson.D{{}}, options)

  // Append to resuls array
  for cursor.Next(ctx){
    product := interfaces.Article{}
    _ = cursor.Decode(&product)
    products = append(products, product)
  }

  return products, nil
}
