package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Plan struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	UserID uint64
	User User
	CategoryID uint8
	Category Category
	Name string `form:"name" binding:"required"`
	Description string `form:"title" binding:"required"`
	Price uint32 `form:"title" binding:"required"`
	Rank uint8
}

func PlanModel(c *gin.Context) {
	c.HTML(200, "plans.html", gin.H{})
}
