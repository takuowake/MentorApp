package user

import (
	"github.com/gin-gonic/gin"
)

func UserPlanManageAction(c *gin.Context){
	c.HTML(200, "user-manage.html", gin.H{})
}