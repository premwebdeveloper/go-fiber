package router

import (
	"web/middleware"
	"web/model/auth"
	employees "web/model/employee"
	"web/model/product"
	users "web/model/user"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {

	router.Get("/create", users.CreateUser)
	router.Get("/users", users.GetUsers)
	router.Get("/user/:id", users.GetUser)
	router.Post("/user", users.SaveUser)
	router.Post("/user-delete/:id", users.DeleteUser)
	router.Put("/user/:id", users.UpdateUser)

	router.Get("/employees", employees.GetEmployees)
	router.Get("/create-employee", employees.CreateEmployee)
	router.Post("/store-employee", employees.StoreEmployee)
	router.Post("/delete-employee/:id", employees.DeleteEmployee)

	router.Post("/login", auth.Login)

	router.Get("/products", product.GetProducts)
	router.Post("product-create", middleware.ValidateToken, product.Create)
	router.Get("product-list", middleware.ValidateToken, product.GetProducts)
	router.Post("product-update", middleware.ValidateToken, product.UpdateProduct)
	router.Post("product-delete", middleware.ValidateToken, product.DeleteProduct)
}
