package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3/log"
)

func GetConnection() (*sql.DB, error) {
	databaseuri := os.Getenv("DATABASE_URI")
	if databaseuri == "" {
		log.Fatal("DATABASE_URI environment variable is not set")
		return nil, fmt.Errorf("DATABASE_URI environment variable is not set")
	}
	db, err := sql.Open("pgx", databaseuri)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return nil, err
	}

	fmt.Println("¡Connected to the database successfully!")

	return db, nil
}
