package model

import "gorm.io/gorm"

//Category displaying the items
type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Items       []Item
}
