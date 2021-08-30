package user

import (
"github.com/gin-gonic/gin"
)

func PaymentManageAction(c *gin.Context){
	c.HTML(200, "payment.html", gin.H{})
}