package model

import "github.com/jinzhu/gorm"

// Categories Model
type Categories struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
}
