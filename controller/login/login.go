package login

import "github.com/gin-gonic/gin"

func LoginAction(c *gin.Context){
	c.HTML(200, "login.html", gin.H{})
}