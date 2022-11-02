package configs

import(
  "fmt"
  "os"
  "context"
  "strconv"
  "time"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/go-redis/redis/v9"
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

  return client
}

// GetCollection get specific collection (table) from mongo client
func GetCollection(collection string) *mongo.Collection{
  col := mongoInstance.Database("buenavida").Collection(collection)
  return col
}

// ### ### ### Postgres ### ### ###
// ConnectToPostgres creates a postgres connecion
func ConnectToPostgres() *sql.DB {
  db, err := sql.Open("postgres", getPostgresURI())

  if err != nil {
    panic("🟥 Unable to create postgres connection 🟥")
  }

  return db
}

// ### ### ### Redis ### ### ###
// ConnectToRedis creates a redis connection
func ConnectToRedis() *redis.Client {
  database, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))

  if err != nil {
    panic("🟥 Unable to parse redis database 🟥")
  }

  client := redis.NewClient(&redis.Options{
    Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
    Password: os.Getenv("REDIS_PASSWORD"),
    DB: database,
  })

  if _, err := client.Ping(context.Background()).Result(); err != nil {
    panic("🟥 Unable to ping redis database 🟥")
  }

  return client
}

// ### ### ### Jwt ### ### ###
// GetJWTSecret get secret from environment
func GetJWTSecret() []byte {
  secret := os.Getenv("JWT_KEY")
  return []byte(secret)
}

// Create instances
var mongoInstance *mongo.Client = ConnectToMongo()
