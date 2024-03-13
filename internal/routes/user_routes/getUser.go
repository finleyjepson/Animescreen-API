package user_routes

import (
	"api/internal/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

// GetUser is a function that returns a user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User

	err := database.DBCon.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(user)
}