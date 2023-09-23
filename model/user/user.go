package model

import (
	connect "web/database"
	users "web/migration/user"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(ctx *fiber.Ctx) error {

	var users []users.User

	connect.DB.Find(&users)

	common := fiber.Map{
		"Title":    "Go Fiber Framework CRUD Operations",
		"Heading":  "My First GO Fiber page.",
		"UserList": "Show All Users",
	}

	return ctx.Render("users/index", fiber.Map{"users": users, "common": common})
}

func CreateUser(ctx *fiber.Ctx) error {
	common := fiber.Map{
		"Title":             "Go Fiber Framework CRUD Operations",
		"Heading":           "My First GO Fiber page.",
		"CreateUserContent": "Create User Form",
	}

	return ctx.Render("users/create", fiber.Map{"common": common})
}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user users.User
	connect.DB.Find(&user, id)
	return ctx.JSON(&user)
}

func SaveUser(ctx *fiber.Ctx) error {
	user := new(users.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	connect.DB.Create(&user)
	ctx.Redirect("/users")
	return ctx.JSON(&user)

}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user users.User
	connect.DB.First(&user, id)
	if user.Email == "" {
		return ctx.Status(500).SendString("User not found !")
	}

	connect.DB.Delete(&user)
	ctx.Redirect("/users")
	return ctx.JSON(&user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user users.User
	connect.DB.First(&user, id)
	if user.Email == "" {
		return ctx.SendString("User not found!")
	}

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	connect.DB.Save(&user)
	return ctx.JSON(&user)
}
