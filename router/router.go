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

	api := router.Group("/api")
	user := api.Group("/users")

	user.Get("/create", users.CreateUser)
	user.Get("/users", users.GetUsers)
	user.Get("/user/:id", users.GetUser)
	user.Post("/user", users.SaveUser)
	user.Post("/user-delete/:id", users.DeleteUser)
	user.Put("/user/:id", users.UpdateUser)

	employee := api.Group("/employees")
	employee.Get("/employees", employees.GetEmployees)
	employee.Get("/create-employee", employees.CreateEmployee)
	employee.Post("/store-employee", employees.StoreEmployee)
	employee.Post("/delete-employee/:id", employees.DeleteEmployee)

	authentic := api.Group("/authentic")
	authentic.Post("/login", auth.Login)

	products := api.Group("/products")
	products.Get("/products", product.GetProducts)
	products.Post("product-create", middleware.ValidateToken, product.Create)
	products.Get("product-list", middleware.ValidateToken, product.GetProducts)
	products.Post("product-update", middleware.ValidateToken, product.UpdateProduct)
	products.Delete("product-delete", middleware.ValidateToken, product.DeleteProduct)
}
