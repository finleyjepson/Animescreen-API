package event_routes

import (
	"api/internal/database"

	"github.com/gofiber/fiber/v2"
)

// Get Event by ID
func GetEventById(c *fiber.Ctx) error {
	id := c.Params("id")
	event := new(Event)
	err := database.DBCon.QueryRow("SELECT * FROM events WHERE id = $1", id).Scan(&event.ID, &event.Title, &event.Description)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(event)
}