package main

import (
	"log"
	"myapp/app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	e := echo.New()
	e.Use(middleware.CORS())
	routes.ECPayRoute(e)
	routes.YoutubeRoute(e)
	routes.TwitchApiRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
