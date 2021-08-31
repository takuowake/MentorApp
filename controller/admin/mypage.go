package admin

import (
"github.com/gin-gonic/gin"
)

func AdminMyPageAction(c *gin.Context){
	c.HTML(200, "admin-mypage.html", gin.H{})
}