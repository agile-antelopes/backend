package main

import (
	"database/sql"
	"fmt"

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
	connStr := "postgres://juansebas053cr__vert__84ydd:V&Hs0cwNR9jj@198.54.124.77:5432/pin21021_Agile_Antelopes"
	db, err := sql.Open("pgx", connStr)
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
