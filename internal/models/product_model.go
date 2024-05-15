package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"name"`
	Price float32 `gorm:"price"`
}
