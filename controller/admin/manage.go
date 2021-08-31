package admin

import (
"github.com/gin-gonic/gin"
)

func AdminPlanManageAction(c *gin.Context){
	c.HTML(200, "admin.html", gin.H{})
}