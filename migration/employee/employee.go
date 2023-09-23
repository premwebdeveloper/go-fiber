package migration

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   int    `json:"mobile"`
	JwtToken string `json:"jwttoken"`
	Address  string `json:"address"`
	Status   string `json:"status"`
}
