package event_routes

import (
	"api/internal/database"

	"github.com/gofiber/fiber/v2"
)

// CreateEvent creates a new event
func CreateEvent(c *fiber.Ctx) error {
	event := new(Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Check if category exists & get category ID
	var CategoryID int
	err := database.DBCon.QueryRow("SELECT id FROM categories WHERE name = $1", event.Category).Scan(&CategoryID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Category does not exist",
		})
	}

	// Insert event into database
	_, err = database.DBCon.Exec("INSERT INTO events (title, description, category_id) VALUES ($1, $2, $3)", event.Title, event.Description, CategoryID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Event created successfully!")
}