package model

import (
	"gorm.io/gorm"
)

// User gives details like UserID,Name,Email,Password,ShoppingList
type User struct {
	gorm.Model
	UserID       string         `gorm:"type:varchar(30);not null"`
	Name         string         `gorm:"not null"`
	Email        string         `gorm:"not null"`
	Password     string         `gorm:"type:varchar(80);not null"`
	ShoppingList []ShoppingList `gorm:"many2many:user_shopping_lists;foreignkey:UserID"`
}
