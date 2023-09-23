package auth

import (
	"time"
	connect "web/database"
	employees "web/migration/employee"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Secret key for signing the JWT token
var jwtSecret = []byte("welcom-to-rginfotech")

func Login(ctx *fiber.Ctx) error {

	employee := new(employees.Employee)

	if err := ctx.BodyParser(&employee); err != nil {
		return err
	}

	existingUser := new(employees.Employee)

	result := connect.DB.Where("username = ?", employee.Username).First(&existingUser)
	if result.Error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username",
		})
	}

	// Validate password (add your own logic here)
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(employee.Password))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = existingUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	// If the user already exists, update the token
	existingUser.JwtToken = tokenString
	connect.DB.Save(&existingUser)

	return ctx.JSON(fiber.Map{
		"status":  200,
		"message": "Login Successfully.",
		"token":   tokenString,
	})

}

func GetTokenForUser(token string) (string, error) {

	var employee employees.Employee

	result := connect.DB.Where("jwt_token = ?", token).First(&employee)

	if result.Error != nil {
		return "", result.Error
	}

	return employee.Username, nil
}
