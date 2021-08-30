package user

import (
	"github.com/gin-gonic/gin"
)

func UserResignAction(c *gin.Context){
	c.HTML(200, "resign.html", gin.H{})
}
