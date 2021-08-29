package controllers

import "github.com/gin-gonic/gin"

func UserAddDisplayAction(c *gin.Context){
	c.HTML(200, "user-plan-add.html", gin.H{})
}

