package main

import (
	"github.com/alexfrunza/servicefinder/api"
	"github.com/labstack/echo/v4"
	"log"

	_ "github.com/alexfrunza/servicefinder/db"
)

func main() {
	e := echo.New()
	err := api.AddRoutes(e)
	if err != nil {
		log.Fatalln(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
