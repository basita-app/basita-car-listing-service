package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

// Connect initializes the database connection
func Connect() error {
	var err error

	// Get Turso database configuration from environment variables
	databaseURL := os.Getenv("TURSO_DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("TURSO_DATABASE_URL environment variable is not set")
	}

	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	if authToken == "" {
		return fmt.Errorf("TURSO_AUTH_TOKEN environment variable is not set")
	}

	// Build connection string for Turso
	// Format: libsql://[database-url]?authToken=[token]
	connectionString := fmt.Sprintf("%s?authToken=%s", databaseURL, authToken)

	// Connect to Turso database
	DB, err = sql.Open("libsql", connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to Turso database: %w", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping Turso database: %w", err)
	}

	log.Println("Turso database connection established")
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return DB
}
