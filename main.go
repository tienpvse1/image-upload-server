package main

import (
	"tienpvse/file"
	"tienpvse/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// create app instance
	app := fiber.New()
	// initialize database connection
	db := InitConnection()
	// recover if app panic
	api := app.Group("api/v1")

	// serve static file at
	app.Static("/", "./upload")

	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	user.InitUserRoute(api, db)
	file.InitImageRoute(api, db)

	app.Listen(":5000")
}
