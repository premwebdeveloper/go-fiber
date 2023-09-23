package model

import (
	"fmt"
	connect "web/database"
	employees "web/migration/employee"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetEmployees(ctx *fiber.Ctx) error {

	var employees []employees.Employee

	connect.DB.Find(&employees)

	common := fiber.Map{
		"Title":        "Go Fiber Framework CRUD Operations",
		"Heading":      "My First GO Fiber page.",
		"EmployeeList": "Show All Employees",
	}

	return ctx.Render("employees/index", fiber.Map{"employees": employees, "common": common})
}

func CreateEmployee(ctx *fiber.Ctx) error {
	common := fiber.Map{
		"Title":                 "Go Fiber Framework CRUD Operations",
		"Heading":               "My First GO Fiber page.",
		"CreateEmployeeContent": "Create Employee Form",
	}

	return ctx.Render("employees/create", fiber.Map{"common": common})
}

func StoreEmployee(ctx *fiber.Ctx) error {
	employee := new(employees.Employee)

	if err := ctx.BodyParser(employee); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	employee.Password = string(hashedPassword)

	connect.DB.Create(&employee)
	ctx.Redirect("employees")
	return ctx.JSON(&employee)
}

func DeleteEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	fmt.Println("ID", id)
	var employee = new(employees.Employee)

	connect.DB.First(&employee, id)
	if employee.Email == "" {
		return ctx.SendString("Employee Not Found!")
	}

	connect.DB.Delete(&employee)
	ctx.Redirect("/employees")
	return ctx.JSON(&employee)
}
