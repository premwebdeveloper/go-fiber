package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"unique" validate:"required,min=5,max=20"`
	Price       float64 `json:"price" gorm:"decimal(10,2)" validate:"required"`
	Description string  `json:"description"`
}
