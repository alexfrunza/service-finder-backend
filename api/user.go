package api

import (
	"fmt"
	"github.com/alexfrunza/servicefinder/db"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type User struct {
	Email       string `json:"email" validate:"required,email,max=300"`
	FirstName   string `json:"firstName" validate:"required,alpha,max=100"`
	LastName    string `json:"lastName" validate:"required,alpha,max=100"`
	PhoneNumber string `json:"phoneNumber" validate:"required,max=50"`
	Password    string `json:"password" validate:"required,max=32,containsany=0123456789"`
	CountryId   int    `json:"countryId" validate:"required,number"`
	CountyId    int    `json:"countyId" validate:"required,number"`
}

func createUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	stmt := `INSERT INTO users(email, first_name, last_name, password, phone_number, country_id, county_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err := db.DB.Exec(stmt, u.Email, u.FirstName, u.LastName, u.Password, u.PhoneNumber, u.CountryId, u.CountyId); err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusCreated, &httpMsg{Message: "User created successfully."})
}

func getUsers(c echo.Context) error {
	rows, err := db.DB.Query(`SELECT email, first_name, last_name, phone_number FROM users;`)
	if err != nil {
		fmt.Println("foo")
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Email, &u.FirstName, &u.LastName, &u.PhoneNumber); err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}
	return c.String(http.StatusOK, "Users")
}
