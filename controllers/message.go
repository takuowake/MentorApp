package controllers

import "github.com/gin-gonic/gin"

func UserMessageDisplayAction(c *gin.Context){
	c.HTML(200, "user-message.html", gin.H{})
}
