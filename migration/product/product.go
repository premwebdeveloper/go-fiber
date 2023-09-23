package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"unique"`
	Price       float64 `json:"price" gorm:"decimal(10,2)"`
	Description string  `json:"description"`
}
