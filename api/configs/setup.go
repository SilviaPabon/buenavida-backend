package configs

import(
  "fmt"
  "context"
  "time"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)


// ConnectToMongo Establish a mongo connection (not instance)
func ConnectToMongo() *mongo.Client{
  // Create client
  client, err := mongo.NewClient(options.Client().ApplyURI(getMongoURI()))

  if err != nil{
    panic("🟥 Unable to create mongo client 🟥")
  }

  // Create context
  ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
  err = client.Connect(ctx)

  if err != nil{
    panic("🟥 Unable to create mongo connection 🟥")
  }

  // Ping database
  err = client.Ping(ctx, nil)

  if err != nil {
    panic("🟥 Unable to get response from dababase 🟥")
  }

  fmt.Println("🟩 Connected to mongo 🟩")
  return client
}

// Create instance
var mongoInstance *mongo.Client = ConnectToMongo()

// GetCollection get specific collection (table) from mongo client
func GetCollection(collection string) *mongo.Collection{
  col := mongoInstance.Database("buenavida").Collection(collection)
  return col
}
