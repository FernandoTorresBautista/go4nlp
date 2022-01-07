package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// render HTML
	// templating engine
	engine := html.New("./views", ".html")

	// reload for changes :for dev
	engine.Reload(true)

	// Init App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// route
	app.Get("/", func(c *fiber.Ctx) error {
		message := "Get - Hello go fiber"

		formName := c.FormValue("name")
		formMessage := c.FormValue("message")
		// return c.Render("index")
		return c.Render("index", fiber.Map{
			"message":     message,
			"formname":    formName,
			"formmessage": formMessage,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		message := "Post - Hello go fiber"

		formName := c.FormValue("name")
		formMessage := c.FormValue("message")
		// return c.Render("index")
		return c.Render("index", fiber.Map{
			"message":     message,
			"formname":    formName,
			"formmessage": formMessage,
		})
	})

	// Listen
	app.Listen(":3000")
}
