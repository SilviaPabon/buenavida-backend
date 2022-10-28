package configs

import(
  "fmt"
  "os"
  "github.com/joho/godotenv"
)

// getPostgresURI return postgres URI created from environment variables
func getPostgresURI() string{
  // Try to get environment variables from system
  user := os.Getenv("PG_USER")
  password := os.Getenv("PG_PASSWORD")
  host := os.Getenv("PG_HOST")
  port := os.Getenv("PG_PORT")

  // Gen environment variables from .env file if needed
  if user == "" || password == "" || host == "" || port == "" {
    err := godotenv.Load()

    if err != nil {
      panic("🟥 Unable to load environment variables 🟥")
    }

    user = os.Getenv("PG_USER")
    password = os.Getenv("PG_PASSWORD")
    host = os.Getenv("PG_HOST")
    port = os.Getenv("PG_PORT")

  }

  URI := fmt.Sprintf("postgres://%s:%s@%s:%s/buenavida?sslmode=disable", user, password, host, port)
  return URI

}

// getMongoURI return mongo URI created from environment variables
func getMongoURI() string {
  // Try to get environment variables from system
  user := os.Getenv("MONGO_USER")
  password := os.Getenv("MONGO_PASSWORD")
  host := os.Getenv("MONGO_HOST")
  port := os.Getenv("MONGO_PORT")

  // Get environment variables from .env file if needed
  if user == "" || password == "" || host == "" || port == "" {
    // Load environment variables from .env file
    err := godotenv.Load()

    if err != nil {
      panic("🟥 Unable to load environment variables 🟥")
    }
    
    user = os.Getenv("MONGO_USER")
    password = os.Getenv("MONGO_PASSWORD")
    host = os.Getenv("MONGO_HOST")
    port = os.Getenv("MONGO_PORT")
  }
  
  // Generate and return URI
  URI := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
  return URI
}
