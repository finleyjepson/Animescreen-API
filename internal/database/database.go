package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)

func ConnectDB() *sql.DB {
	connStr := os.Getenv("DB_URI")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")

	if err != nil {
		log.Fatal(err)
	}

	return db
}