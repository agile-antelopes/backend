package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Question struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Country  string `json:"country"`
	UserName string `json:"username"`
}

func main() {
	databaseuri := os.Getenv("DATABASE_URI")
	if databaseuri == "" {
		log.Fatal("DATABASE_URI environment variable is not set")
		return
	}
	db, err := sql.Open("pgx", databaseuri)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {

	}

	fmt.Println("¡Conexión a PostgreSQL establecida con éxito!")
	app := fiber.New()

	app.Post("/question/save", func(c fiber.Ctx) error {
		var question Question
		err := c.Bind().Body(&question)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		fmt.Println("Save into database...")

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": "Saved question",
		})
	})

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
