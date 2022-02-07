package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func renderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":             "Welcome to Liam's World",
		"India-Description": "Let's visit India",
		"Dubai-Description": "Let's visit Dubai",
		"USA-Description":   "Let's visit USA",
	}, "layouts/main")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", renderIndex)
}

func main() {
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)
	app.Listen(":3000")
}
