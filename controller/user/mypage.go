package user

import (
"github.com/gin-gonic/gin"
)

func MyPageAction(c *gin.Context){
	c.HTML(200, "mypage.html", gin.H{})
}