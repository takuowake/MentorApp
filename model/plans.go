package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Plan struct {
	gorm.Model
	ID string `gorm:"primaryKey"`
	UserID int
	User User
	CategoryID int
	Category Category
	name string `form:"name" binding:"required"`
	description string `form:"title" binding:"required"`
	price int `form:"title" binding:"required"`
	rank int
}

type Category struct {
	ID int
	name string
	states string
}

func PlanModel(c *gin.Context) {
	c.HTML(200, "plans.html", gin.H{})
}
