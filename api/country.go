package api

import (
	"github.com/alexfrunza/servicefinder/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

type country struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type resCountries struct {
	Count     int       `json:"count"`
	Countries []country `json:"countries"`
}

func getCountries(c echo.Context) error {
	var counter int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM countries;`).Scan(&counter)
	if err != nil {
		return err
	}

	rows, err := db.DB.Query(`SELECT id, name FROM countries;`)
	if err != nil {
		return err
	}
	defer rows.Close()

	countries := make([]country, 0, counter)
	for rows.Next() {
		var c country
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return err
		}
		countries = append(countries, c)
	}

	return c.JSON(http.StatusOK, &resCountries{Count: counter, Countries: countries})
}
