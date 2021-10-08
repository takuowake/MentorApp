package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Plan struct {
	gorm.Model
	title string
	content string
}

func PlanModel(c *gin.Context) {
	c.HTML(200, "plans.html", gin.H{})
}


func CreatePlan(title string, content string) []error {
	dbInit()
	db := gormConnect()
	defer db.Close()
	return nil
}
