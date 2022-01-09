package model

import "gorm.io/gorm"

//ShoppingList gives the title for the user ShoppingList
type ShoppingList struct {
	gorm.Model
	Title string `gorm:"not null"`
	Items []Item
	Users []User `gorm:"many2many:user_shopping_lists;association_foreignkey:UserID"`
}
