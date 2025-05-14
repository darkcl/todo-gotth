package main

import (
	"fmt"

	"todo-gotth/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	InitDotEnv()

	e := echo.New()

	// Set up the asset server
	routes.SetupAssetsRoutes(e)
	routes.SetupAPIRoutes(e)
	routes.SetupPagesRoutes(e)

	// Use logger for echo
	e.Use(middleware.Logger())


	fmt.Println("Server is running on http://localhost:8090")
	e.Logger.Fatal(e.Start(":8090"))
}

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
