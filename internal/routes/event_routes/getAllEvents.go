package event_routes

import (
	"api/internal/database"

	"github.com/gofiber/fiber/v2"
)

// Get Event by ID
func GetAllEvents(c *fiber.Ctx) error {
	events := []Event{}

	rows, err := database.DBCon.Query("SELECT * FROM events")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Title, &event.Description)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		events = append(events, event)
	}

	return c.JSON(events)
}