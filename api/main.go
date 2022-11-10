package main

import (
	"fmt"
	"os"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/SilviaPabon/buenavida-backend/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// ### ### ###
	// Create mongo database connection
	configs.ConnectToMongo()

	//db := configs.ConnectToPostgres()

	// ### ### ###
	// Echo setup
	e := echo.New()

	// Change this to "http://localhost:5173" for local tests
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"hhttps://splendid-piroshki-dade21.netlify.app"},
		AllowCredentials: true,
	}))

	// Start routes
	routes.SetupProductsRoutes(e)
	routes.SetupUserRoutes(e)
	routes.SetupSessionRoutes(e)
	routes.SetupCartRoutes(e)

	// ### ### ###
	// Configure port

	port := os.Getenv("PORT")

	if port == "" {
		port = "3030"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
