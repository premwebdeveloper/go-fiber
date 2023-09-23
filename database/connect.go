package connect

import (
	"fmt"
	"log"
	employees "web/migration/employee"
	"web/migration/product"
	users "web/migration/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Prem@2402@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("DB connection failed!")
	}

	err = DB.AutoMigrate(&users.User{}, &employees.Employee{}, &product.Product{})
	if err != nil {
		log.Fatal(err)
	}
}
