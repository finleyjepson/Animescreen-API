package user_routes

import (
	"api/internal/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	fetchedUsers, err := database.DBCon.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var data []User

	for fetchedUsers.Next() {
		var user User
		err := fetchedUsers.Scan(&user.ID, &user.Username, &user.Password, &user.IsOrganiser, &user.IsAdmin, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, user)
	}

	return c.JSON(fiber.Map{
		"users": data,
	})
}