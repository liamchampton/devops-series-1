package main

// imports used by the project
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

/* renderIndex - renders the index.hbs page
** Input = fiber context
** Output = passes map to layouts/main to be rendered / error
**/
func renderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":             "Welcome to Liam's World",
		"India-Description": "Let's visit India",
		"Dubai-Description": "Let's visit Dubai",
		"USA-Description":   "Let's visit USA",
	}, "layouts/main")
}

/* setupRoutes - abstracted function for all application routes
** Input = pointer to fiber App struct
** Output = null
**/
func setupRoutes(app *fiber.App) {
	app.Get("/", renderIndex)
}

/* main - main function executing the server and handlebars templating
** Input = null
** Output = null / listens on port 3000 for the app server
**/
func main() {
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)
	app.Listen(":3000")
}
