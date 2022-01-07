package main

import (
	"github.com/gofiber/fiber/v2"
)

// Method 2: func
func home(c *fiber.Ctx) error {
	return x.SendString("Hello Go Fiber")
}

func main() {
	// Init App
	app := fiber.New()

	// Route
	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello Go Fiber")
	})
	// method 2:
	app.Get("/home", home)
	// app.Post()

	// Params
	// localhost:3000/api/name
	app.Get("/api/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		// fmt.Println(name)
		return c.SendString("Hello " + name)
	})

	// Listen
	app.Listen(":3000")
}
