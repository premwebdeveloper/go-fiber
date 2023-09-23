package main

import (
	connect "web/database"
	"web/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title":             "Go Fiber Framework CRUD Operations",
		"Heading":           "My First GO Fiber page.",
		"CreateUserContent": "Create User Form",
		"UserList":          "Show All Users",
		"EditUserContent":   "Edit User Form",
	})
}

func main() {

	connect.InitialMigration()

	// Set up the HTML template engine
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.Initialize(app)

	app.Get("/", Index)

	app.Listen(":4000")
}
