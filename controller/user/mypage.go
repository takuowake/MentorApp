package user

import (
"github.com/gin-gonic/gin"
)

func UserMyPageAction(c *gin.Context){
	c.HTML(200, "user-mypage.html", gin.H{})
}