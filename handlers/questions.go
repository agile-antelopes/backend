package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// --- DTOs (Data Transfer Objects) ---

type CountryPayload struct {
	CountryCode string            `json:"country_code"`
	CountryName string            `json:"country_name"`
	Details     map[string]string `json:"details"` // El JSON dinámico
}

type ResponsePayload struct {
	TopicTagID int    `json:"topic_tag_id"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}

type InterviewPayload struct {
	InterviewerName string            `json:"interviewer_name"`
	IntervieweeName string            `json:"interviewee_name"`
	CountryID       string            `json:"country_id"`
	Responses       []ResponsePayload `json:"responses"`
}

// --- DTOs para las respuestas ---
type CountryResponse struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Facts       string `json:"facts"`
}

type TopicResponse struct {
	TopicTagID int    `json:"topic_tag_id"`
	TopicTag   string `json:"topic_tag"`
}

type QuestionResponse struct {
	ID           int    `json:"id"`
	TopicTagID   int    `json:"topic_tag_id"`
	QuestionText string `json:"question_text"`
}

// --- DTO para la pregunta ---
type QuestionPayload struct {
	TopicTagID   int    `json:"topic_tag_id"`
	QuestionText string `json:"question_text"`
}

// CreateQuestion recibe el Formulario 3 y guarda en la tabla question_template
func CreateQuestion(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var payload QuestionPayload

		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Validamos que no vengan campos vacíos
		if payload.TopicTagID == 0 || payload.QuestionText == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "topic_tag_id and question_text are required"})
		}

		var newID int
		// Usamos RETURNING para obtener el ID que PostgreSQL genera automáticamente
		query := `
			INSERT INTO worldloom.question_template (topic_tag_id, question_text) 
			VALUES ($1, $2) 
			RETURNING question_id
		`
		err := db.QueryRow(query, payload.TopicTagID, payload.QuestionText).Scan(&newID)
		if err != nil {
			log.Errorf("Error inserting question: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save question"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success":     true,
			"message":     "Question template created",
			"question_id": newID,
		})
	}
}

// --- HANDLERS GET ---

// GetCountries extrae todos los países del esquema worldloom
func GetCountries(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		rows, err := db.Query(`SELECT country_code, country_name, facts FROM worldloom.country ORDER BY country_name ASC`)
		if err != nil {
			log.Errorf("Error fetching countries: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch countries"})
		}
		defer rows.Close()

		countries := []CountryResponse{} // Inicializamos el slice vacío
		for rows.Next() {
			var country CountryResponse
			if err := rows.Scan(&country.CountryCode, &country.CountryName, &country.Facts); err != nil {
				log.Errorf("Error scanning country row: %v", err)
				continue
			}
			countries = append(countries, country)
		}

		return c.JSON(countries)
	}
}

// GetTopics extrae todas las categorías
func GetTopics(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		rows, err := db.Query(`SELECT topic_tag_id, topic_tag FROM worldloom."topic_tag-A" ORDER BY topic_tag_id ASC`)
		if err != nil {
			log.Errorf("Error fetching topics: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch topics"})
		}
		defer rows.Close()

		topics := []TopicResponse{}
		for rows.Next() {
			var topic TopicResponse
			if err := rows.Scan(&topic.TopicTagID, &topic.TopicTag); err != nil {
				log.Errorf("Error scanning topic row: %v", err)
				continue
			}
			topics = append(topics, topic)
		}

		return c.JSON(topics)
	}
}

// GetQuestions extrae el banco de preguntas (plantillas)
func GetQuestions(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		rows, err := db.Query(`SELECT question_id, topic_tag_id, question_text FROM worldloom.question_template ORDER BY topic_tag_id ASC`)
		if err != nil {
			log.Errorf("Error fetching questions: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch questions"})
		}
		defer rows.Close()

		questions := []QuestionResponse{}
		for rows.Next() {
			var q QuestionResponse
			if err := rows.Scan(&q.ID, &q.TopicTagID, &q.QuestionText); err != nil {
				log.Errorf("Error scanning question row: %v", err)
				continue
			}
			questions = append(questions, q)
		}

		return c.JSON(questions)
	}
}

// --- HANDLERS ---

// CreateCountry recibe el Formulario 2 y guarda en la tabla `country`
func CreateCountry(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var payload CountryPayload

		// Parseamos el JSON del frontend
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Convertimos el map de details a un string JSON para la columna 'facts'
		factsJSON, err := json.Marshal(payload.Details)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process facts"})
		}

		// Insertamos en la BD
		query := `INSERT INTO worldloom."country" (country_code, country_name, facts) VALUES ($1, $2, $3)`
		_, err = db.Exec(query, payload.CountryCode, payload.CountryName, string(factsJSON))
		if err != nil {
			log.Errorf("Error inserting country: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save country"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "message": "Country created"})
	}
}

// CreateInterview recibe el Formulario 1 y guarda en `interview-A` y `response-A`
func CreateInterview(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var payload InterviewPayload

		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Iniciamos una transacción. Todo o nada.
		tx, err := db.Begin()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to start transaction"})
		}
		// Defer para hacer Rollback si la función termina con un panic o un error no manejado
		defer tx.Rollback()

		// 1. Insertar la Entrevista (Usamos CURRENT_DATE para la fecha y obtenemos el ID generado)
		var interviewID int
		interviewQuery := `
			INSERT INTO worldloom."interview-A" (interviewer_name, interviewee_name, country_id, date) 
			VALUES ($1, $2, $3, CURRENT_DATE) 
			RETURNING interview_id
		`
		err = tx.QueryRow(interviewQuery, payload.InterviewerName, payload.IntervieweeName, payload.CountryID).Scan(&interviewID)
		if err != nil {
			log.Errorf("Error inserting interview: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save interview"})
		}

		// 2. Insertar cada una de las respuestas iterando sobre el array
		responseQuery := `
			INSERT INTO worldloom."response-A" (interview_id, country_id, topic_tag_id, question, answer) 
			VALUES ($1, $2, $3, $4, $5)
		`
		for _, resp := range payload.Responses {
			// Solo insertamos si realmente hay una respuesta
			if resp.Answer != "" {
				_, err = tx.Exec(responseQuery, interviewID, payload.CountryID, resp.TopicTagID, resp.Question, resp.Answer)
				if err != nil {
					log.Errorf("Error inserting response: %v", err)
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save responses"})
				}
			}
		}

		// 3. Si todo salió bien, hacemos Commit a la base de datos
		if err := tx.Commit(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Transaction failed"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success":      true,
			"message":      "Interview and responses saved successfully",
			"interview_id": interviewID,
		})
	}
}
