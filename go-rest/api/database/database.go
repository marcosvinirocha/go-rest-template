package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InitDB initializes the database connection
func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	return db
}