package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// render HTML templating engine
	engine := html.New("./views", ".html")

	// Init App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// serving static files
	app.Static("/", "./public")

	// route
	app.Get("/", func(c *fiber.Ctx) error {
		message := "Get - Hello go fiber"
		return c.Render("index", fiber.Map{
			"message": message,
		})
	})

	// Listen
	app.Listen(":3000")
}
