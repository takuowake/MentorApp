package order

import "github.com/jinzhu/gorm"

type Orders struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	
}
