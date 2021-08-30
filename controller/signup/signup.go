package signup

import "github.com/gin-gonic/gin"

func SignUpAction(c *gin.Context){
	c.HTML(200, "signup.html", gin.H{})
}