package plan

import (
	"github.com/gin-gonic/gin"
)

func ListDisplayAction(c *gin.Context){
	c.HTML(200, "list.html", gin.H{})
}