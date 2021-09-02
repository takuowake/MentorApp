package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Userモデルの宣言
type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

func LoginAction(c *gin.Context){
	c.HTML(200, "login.html", gin.H{})
}