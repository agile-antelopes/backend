package main

import (
	"backend/database"
	"backend/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{"Origin, Content-Type, Accept"},
	}))

	app.Get("/countries", handlers.GetCountries(db))
	app.Get("/topics", handlers.GetTopics(db))
	app.Get("/questions", handlers.GetQuestions(db))

	app.Post("/countries", handlers.CreateCountry(db))
	app.Post("/interviews", handlers.CreateInterview(db))
	app.Post("/topics", handlers.CreateTopics(db))
	app.Post("/questions", handlers.CreateQuestion(db))

	err = app.Listen(":3001")
	if err != nil {
		log.Fatal(err)
	}
}
