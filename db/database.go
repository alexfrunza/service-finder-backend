package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func init() {
	var err error

	connStr, ok := os.LookupEnv("POSTGRESQL_URL")
	if !ok {
		log.Fatalln("You must set in environment variables POSTGRESQL_URL!")
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected to the database!")
}
