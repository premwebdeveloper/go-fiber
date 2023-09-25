package product

import (
	"errors"
	"fmt"
	"log"
	connect "web/database"
	"web/migration/product"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type DeleteRequest struct {
	ID int `json:"id"`
}

var validate *validator.Validate

func GetProducts(ctx *fiber.Ctx) error {
	var product []product.Product
	connect.DB.Find(&product)
	return ctx.JSON(&product)
}

func Create(ctx *fiber.Ctx) error {

	product := new(product.Product)

	// username := ctx.Locals("username").(string)
	// return ctx.SendString("Hello, " + username + "!")

	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	validate = validator.New()

	// Validate form data
	if err := validate.Struct(product); err != nil {
		// Map field-level errors to a map
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("%s is %s , Required %s", err.Field(), err.Tag(), err.Value())
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	result := connect.DB.Create(&product)

	if result.Error != nil {
		fmt.Println("Error occurred during INSERT:", result.Error)

		// Check if the error is due to a unique constraint violation
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return ctx.Status(fiber.StatusConflict).SendString("Duplicate entry!")
		}
		log.Println(result.Error)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error occurred")

	} else {
		fmt.Printf("Rows affected: %d\n", result.RowsAffected)
	}

	return ctx.JSON(&product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product product.Product
	connect.DB.Find(&product, id)
	if product.ID == 0 {
		ctx.SendString("Product Not Found!")
	}

	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	connect.DB.Save(&product)
	return ctx.JSON(&product)
}

func DeleteProduct(ctx *fiber.Ctx) error {

	request := new(DeleteRequest)
	var product product.Product

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	result := connect.DB.First(&product, request)

	if result.Error != nil {
		return ctx.SendString("Error :" + result.Error.Error())
	}

	connect.DB.Delete(&product)
	return ctx.SendString("Product " + product.Name + " deleted successfully.")
}
