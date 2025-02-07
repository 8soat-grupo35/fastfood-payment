package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID uint32 `gorm:"not null"`
	Status  string `gorm:"not null"`
}