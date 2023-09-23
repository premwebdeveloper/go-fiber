package product

import (
	"errors"
	"fmt"
	"log"
	connect "web/database"
	"web/migration/product"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

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
	id := ctx.Params("id")

	return ctx.JSON(id)
	// var product product.Product
	// result := connect.DB.Find(&product, id)

	// if result.Error != nil {
	// if result.Error == gorm.ErrRecordNotFound {
	// 	fmt.Println("Record not found")
	// } else {
	// 	fmt.Println("Error occurred:", result.Error)
	// }
	// }

	// connect.DB.Delete(&product)
	// return ctx.SendString("Product " + product.Name + " deleted successfully.")
}
