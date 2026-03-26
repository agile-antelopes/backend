package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type CountryDetails struct {
	CountryCode string            `json:"country_code"`
	CountryName string            `json:"country_name"`
	Facts       map[string]string `json:"facts"`
	Interviews  []UserExperience  `json:"experiences"`
}

type UserExperience struct {
	Interviewee string             `json:"interviewee"`
	Interviewer string             `json:"interviewer"`
	Date        string             `json:"date"`
	Responses   []DetailedResponse `json:"responses"`
}

type DetailedResponse struct {
	Topic    string `json:"topic"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func GetCountryFullDetails(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		countryCode := c.Params("code")

		// 1. Obtener información básica del país
		var country CountryDetails
		var factsRaw string
		queryCountry := `SELECT country_code, country_name, facts FROM worldloom.country WHERE country_code = $1`

		err := db.QueryRow(queryCountry, countryCode).Scan(&country.CountryCode, &country.CountryName, &factsRaw)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(404).JSON(fiber.Map{"error": "Country not found"})
			}
			return c.Status(500).JSON(fiber.Map{"error": "Database error"})
		}

		// Convertir el string JSON de facts a un map de Go
		json.Unmarshal([]byte(factsRaw), &country.Facts)

		// 2. Obtener todas las entrevistas y opiniones de este país
		queryInterviews := `
			SELECT 
				i.interviewee_name, i.interviewer_name, i.date, 
				t.topic_tag, r.question, r.answer
			FROM worldloom."interview-A" i
			JOIN worldloom."response-A" r ON i.interview_id = r.interview_id
			JOIN worldloom."topic_tag-A" t ON r.topic_tag_id = t.topic_tag_id
			WHERE i.country_id = $1
			ORDER BY i.date DESC
		`
		rows, err := db.Query(queryInterviews, countryCode)
		if err != nil {
			log.Errorf("Error fetching experiences: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch experiences"})
		}
		defer rows.Close()

		// Agrupar respuestas por entrevistado
		experienceMap := make(map[string]*UserExperience)
		for rows.Next() {
			var interviewee, interviewer, date, topic, question, answer string
			rows.Scan(&interviewee, &interviewer, &date, &topic, &question, &answer)

			if _, ok := experienceMap[interviewee]; !ok {
				experienceMap[interviewee] = &UserExperience{
					Interviewee: interviewee,
					Interviewer: interviewer,
					Date:        date,
					Responses:   []DetailedResponse{},
				}
			}
			experienceMap[interviewee].Responses = append(experienceMap[interviewee].Responses, DetailedResponse{
				Topic:    topic,
				Question: question,
				Answer:   answer,
			})
		}

		// Convertir el mapa a un slice para el JSON final
		for _, exp := range experienceMap {
			country.Interviews = append(country.Interviews, *exp)
		}

		return c.JSON(country)
	}
}
