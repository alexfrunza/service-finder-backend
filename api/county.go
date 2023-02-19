package api

import (
	"fmt"
	"github.com/alexfrunza/servicefinder/db"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type county struct {
	Name      string `json:"name"`
	Id        int    `json:"id"`
	CountryId int    `json:"countryId"`
}

type resCounties struct {
	Count    int      `json:"count"`
	Counties []county `json:"counties"`
}

func getCounties(c echo.Context) error {
	countQuery := `SELECT COUNT(*) FROM counties `
	rowsQuery := `SELECT id, name, country_id  FROM counties `

	if c.QueryParam("countryId") != "" {
		countryId, err := strconv.Atoi(c.QueryParam("countryId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, &resBodyError{Errors: []apiError{{Message: "This field must be an integer", Param: "countryId"}}})
		}

		countQuery = countQuery + fmt.Sprintf(`WHERE country_id=%d `, countryId)
		rowsQuery = rowsQuery + fmt.Sprintf(`WHERE country_id=%d `, countryId)
	}

	var counter int
	err := db.DB.QueryRow(countQuery).Scan(&counter)
	if err != nil {
		return err
	}

	rows, err := db.DB.Query(rowsQuery)
	if err != nil {
		return err
	}

	counties := make([]county, 0, counter)
	for rows.Next() {
		var c county
		if err := rows.Scan(&c.Id, &c.Name, &c.CountryId); err != nil {
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
