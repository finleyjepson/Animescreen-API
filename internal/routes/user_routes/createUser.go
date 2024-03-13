package user_routes

import (
	"api/internal/database"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var u *User = new(User)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	switch {
		case u.Username == "":
			return c.Status(400).SendString("No username provided")
		case u.Password == "":
			return c.Status(400).SendString("No password provided")
		case len(u.Password) <= 8 || len(u.Password) >= 50:
			return c.Status(400).SendString("Password must be between 8 and 50 characters long")
		case len(u.Username) <= 3 || len(u.Username) >= 20:
			return c.Status(400).SendString("Username must be between 3 & 20 characters long")
	}

	var userID int
	err := database.DBCon.QueryRow("SELECT id FROM users WHERE username = $1", u.Username).Scan(&userID)
	if err != sql.ErrNoRows {
		return c.Status(400).SendString("Username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.DBCon.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, string(hashedPassword))
	if err != nil {
		return c.Status(500).SendString("Error creating user")
	}

	return c.Status(201).SendString("User created")
}