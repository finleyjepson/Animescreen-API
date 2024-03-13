package routes

import (
	"api/internal/routes/category_routes"
	"api/internal/routes/event_routes"
	"api/internal/routes/user_routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewServer() error {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	// User routes
	user := api.Group("/user")
	user.Get("/", user_routes.GetAllUsers)
	user.Post("/", user_routes.CreateUser)
	user.Get("/:id", user_routes.GetUser)
	user.Put("/:id", user_routes.EditUser)
	user.Post("/login", user_routes.Login)

	// Category routes
	category := api.Group("/category")
	category.Get("/", category_routes.GetCategories)

	// Event routes
	event := api.Group("/event")
	event.Post("/", event_routes.CreateEvent)
	event.Get("/:id", event_routes.GetEventById)
	event.Get("/", event_routes.GetAllEvents)

	app.Listen(":" + os.Getenv("PORT"))

	return nil
}