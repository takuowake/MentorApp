package model

import "github.com/jinzhu/gorm"

// Category Model
type Category struct {
	gorm.Model
	ID uint8 `json:"id" gorm:"primaryKey; autoIncrement"`
	name string `gorm:"unique;not null"`
	states string
}
