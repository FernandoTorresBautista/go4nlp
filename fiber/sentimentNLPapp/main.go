package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

type SentimentDetails struct {
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Neutral  float64 `json:"neutral"`
	Compound float64 `json:"compound"`
}

func main() {

	// render HTML templating engine
	engine := html.New("./views", ".html")

	//reload
	engine.Reload(true)

	// Init App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// serving static files
	app.Static("/", "./public")

	// route
	app.Get("/", func(c *fiber.Ctx) error {
		initMessage := "Get - Hello go fiber"
		return c.Render("index", fiber.Map{
			"initMessage": initMessage,
		})
	})

	//
	app.Post("/", func(c *fiber.Ctx) error {
		initMessage := "Get - Hello go fiber"
		message := c.FormValue("message")
		sentimentResults := sentimentize(message)
		return c.Render("index", fiber.Map{
			"initMessage":      initMessage,
			"originalMsg":      message,
			"sentimentDetails": sentimentResults,
		})
	})

	// Listen
	app.Listen(":3000")
}

// Functions for processing
func sentimentize(docx string) SentimentDetails {
	// parse
	parsedtext := sentitext.Parse(docx, lexicon.DefaultLexicon)
	// Process
	results := sentitext.PolarityScore(parsedtext)
	sentimentScore := SentimentDetails{
		Positive: results.Positive,
		Negative: results.Negative,
		Neutral:  results.Neutral,
		Compound: results.Compound,
	}
	return sentimentScore
}
