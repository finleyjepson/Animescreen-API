package user_routes

import (
	"api/internal/database"

	"github.com/gofiber/fiber/v2"
)

type UserEdit struct {
	Password string
}

func EditUser(c *fiber.Ctx) error {
	u := new(UserEdit)

	id := c.Params("id")

	if err := c.BodyParser(u); err != nil {
		return err
	}

	if id == "" {
		return c.Status(400).SendString("No ID provided")
	} else if u.Password == "" {
		return c.Status(400).SendString("No password provided")
	} else {
		_, err := database.DBCon.Exec("UPDATE users SET password = $1 WHERE id = $2", u.Password, id)
		if err != nil {
			return c.Status(500).SendString("Error updating user")
		}
	}

	return c.Status(200).SendString("User updated")
}