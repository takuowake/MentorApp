package model

import "github.com/jinzhu/gorm"

// Category Model
type Category struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	states string
}
