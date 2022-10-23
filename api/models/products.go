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
  options := new(options.FindOptions)
  options.SetSkip(int64(page -1) * 5) // 5 is the limit
  options.SetLimit(5)

  // Make query
  var results []bson.M
  var products []interfaces.Article
  cursor, _ := productsCollection.Find(ctx, bson.D{{}}, options)

  // Convert to bson and struct
  if err := cursor.All(ctx, &results); err != nil{
    return products, err
  }

  // Append to resuls array
  for _, product := range results {
    bsonBytes, _ := bson.Marshal(product)

    var product interfaces.Article
    bson.Unmarshal(bsonBytes, &product)
    products = append(products, product)
  }

  return products, nil
}
