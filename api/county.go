package api

import (
	"github.com/alexfrunza/servicefinder/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

type county struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type resCounties struct {
	Count    int      `json:"count"`
	Counties []county `json:"counties"`
}

func getCounties(c echo.Context) error {
	var counter int
	err := db.DB.QueryRow(`SELECT COUNT(*) FROM counties;`).Scan(&counter)
	if err != nil {
		return err
	}

	rows, err := db.DB.Query(`SELECT id, name FROM counties;`)
	if err != nil {
		return err
	}

	counties := make([]county, 0, counter)
	for rows.Next() {
		var c county
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return err
		}
		counties = append(counties, c)
	}

	err = rows.Close()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &resCounties{Count: counter, Counties: counties})
}
