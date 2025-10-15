package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost/mydb?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "default_secret_key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// InitDatabase initializes the database connection based on the configuration
func InitDatabase(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("Failed to ping database: " + err.Error())
	}

	return db
}