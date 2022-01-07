package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Init App
	app := fiber.New()

	// Route
	app.Get("/api/:name", func(c *fiber.Ctx) error {
		fname := c.Params("name")
		return c.SendString(fname)
	})

	// query params /api/?text
	app.Get("/api/:text?", func(c *fiber.Ctx) error {
		text := c.Query("text")
		newmsg := strings.ToUpper(text)
		// return c.SendString(text)
		return c.JSON(fiber.Map{
			"original": text,
			"modified": newmsg,
		})
	})
	// Listen
	app.Listen(":3000")
}
