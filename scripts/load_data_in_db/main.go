package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type country struct {
	Name string `json:"name"`
}

type county struct {
	Name string `json:"nume"`
}

func loadCountries(db *sql.DB, fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	var countries []country
	err = json.Unmarshal(data, &countries)
	if err != nil {
		return err
	}

	insertStmt := `INSERT INTO countries (name) values ($1)`
	for _, c := range countries {
		if _, err := db.Exec(insertStmt, c.Name); err != nil {
			return err
		}
	}

	return nil
}

func loadCounties(db *sql.DB, fileName string, countryId int) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	var counties []county
	err = json.Unmarshal(data, &counties)
	if err != nil {
		return err
	}

	insertStmt := `INSERT INTO COUNTIES (name, country_id) values ($1, $2)`
	for _, c := range counties {
		if _, err := db.Exec(insertStmt, c.Name, countryId); err != nil {
			return err
		}
	}

	return nil
}

func getCountryIdByName(db *sql.DB, name string) (int, error) {
	var id int
	queryStmt := `SELECT id FROM countries WHERE name = $1`
	if err := db.QueryRow(queryStmt, name).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func main() {
	log.SetFlags(0)

	connStr, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatalln("You must set DATABASE_URL!")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = loadCountries(db, "countries.json")
	if err != nil {
		log.Fatalln(err)
	}

	romaniaId, err := getCountryIdByName(db, "Romania")
	if err != nil {
		log.Fatalln(err)
	}

	err = loadCounties(db, "counties.json", romaniaId)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Countries and counties loaded successfully into the database!")
}
