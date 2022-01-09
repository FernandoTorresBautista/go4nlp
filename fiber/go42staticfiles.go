package main

import "github.com/gofiber/fiber/v2"

func main() {
	// init app
	app := fiber.New()

	// serve static files
	// localhost:3000/image.png
	app.Static("/", "./myfiles")

	// listen
	app.Listen(":3000")
}
