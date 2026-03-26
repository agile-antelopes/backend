package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type BulkTopicPayload struct {
	Topics []string `json:"topics"`
}

func CreateTopics(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var payload BulkTopicPayload

		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if len(payload.Topics) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No topics provided"})
		}

		tx, err := db.Begin()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to start transaction"})
		}
		defer tx.Rollback()

		query := `INSERT INTO worldloom."topic_tag-A" (topic_tag) VALUES ($1)`

		for _, topic := range payload.Topics {
			_, err = tx.Exec(query, topic)
			if err != nil {
				log.Errorf("Error inserting topic '%s': %v", topic, err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Failed to insert topic",
					"details": err.Error(),
				})
			}
		}

		if err := tx.Commit(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Transaction failed"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"message": "All topics inserted successfully",
		})
	}
}
