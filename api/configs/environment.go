package configs

import(
  "fmt"
  "os"
  "github.com/joho/godotenv"
)

// getMongoURI return mongo URI created from environment variables
func getMongoURI() string {
  // Load environment variables
  err := godotenv.Load()

  if err != nil {
    panic("🟥 Unable to load environment variables 🟥")
  }

  // Get environment variables
  user := os.Getenv("MONGO_USER")
  password := os.Getenv("MONGO_PASSWORD")
  host := os.Getenv("MONGO_HOST")
  port := os.Getenv("MONGO_PORT")

  // Generate and return URI
  URI := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
  return URI
}
