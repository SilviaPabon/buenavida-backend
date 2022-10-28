package configs

import(
  "fmt"
  "context"
  "time"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "database/sql"
  _ "github.com/lib/pq"
)

// ### ### ### MONGO ### ### ###
// ConnectToMongo Establish a mongo connection (not instance)
func ConnectToMongo() *mongo.Client{
  // Create client
  client, err := mongo.NewClient(options.Client().ApplyURI(getMongoURI()))

  if err != nil{
    panic("🟥 Unable to create mongo client 🟥")
  }

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()
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

// GetCollection get specific collection (table) from mongo client
func GetCollection(collection string) *mongo.Collection{
  col := mongoInstance.Database("buenavida").Collection(collection)
  return col
}

// ### ### ### Postgres ### ### ###\
// ConnectToPostgres creates a postgres connecion
func ConnectToPostgres() *sql.DB {
  db, err := sql.Open("postgres", getPostgresURI())

  if err != nil {
    panic("🟥 Unable to create postgres connection 🟥")
  }

  return db
}

// Create instances
var mongoInstance *mongo.Client = ConnectToMongo()
