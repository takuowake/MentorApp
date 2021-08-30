package user

import (
"github.com/gin-gonic/gin"
)

func UserMessageAction(c *gin.Context){
	c.HTML(200, "message.html", gin.H{})
}