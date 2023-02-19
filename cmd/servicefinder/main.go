package main

import (
	"fmt"
	"github.com/alexfrunza/servicefinder/api"
	"github.com/labstack/echo/v4"
	"log"
	"os"

	_ "github.com/alexfrunza/servicefinder/db"
)

func main() {
	e := echo.New()
	err := api.AddRoutes(e)
	if err != nil {
		log.Fatalln(err)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
