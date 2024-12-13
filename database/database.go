package database

import (
	"database/sql"
	"log"
)

// database connection vars:
const dbUser = "postgres"
const dbPassword = "admin"
const dbname = "gtrade"
const sslmode = "disable"

func DatabaseConnection() *sql.DB {
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbname + " sslmode=" + sslmode + " "
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
