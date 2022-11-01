package configs

import(
  "fmt"
  "os"
  "github.com/joho/godotenv"
)

func InitEnv(){
  // Try to get environment variables from system
  var done = true 
  var secrets = []string{ os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("JWT_KEY") }

  // Check if all the required variables were loaded
  for _, value := range(secrets) {
    if value == ""{
      done = false
      err := godotenv.Load()

      if(err != nil){
	panic("🟥 Unable to load environment variables FROM .env FILE 🟥")
      }

      break
    }
  }
  
  if !done {
    // Reload secrets
    secrets = []string{ os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("JWT_KEY") }

    for _, value := range(secrets) {
      if value == ""{
	panic("🟥 Unable to load ALL the required environment variables 🟥")
      }
    }
  }

}

// getPostgresURI return postgres URI created from environment variables
func getPostgresURI() string{
  InitEnv()

  user := os.Getenv("PG_USER")
  password := os.Getenv("PG_PASSWORD")
  host := os.Getenv("PG_HOST")
  port := os.Getenv("PG_PORT")

  URI := fmt.Sprintf("postgres://%s:%s@%s:%s/buenavida?sslmode=disable", user, password, host, port)
  return URI

}

// getMongoURI return mongo URI created from environment variables
func getMongoURI() string {
  InitEnv() 

  user := os.Getenv("MONGO_USER")
  password := os.Getenv("MONGO_PASSWORD")
  host := os.Getenv("MONGO_HOST")
  port := os.Getenv("MONGO_PORT")
  
  // Generate and return URI
  URI := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
  return URI
}
