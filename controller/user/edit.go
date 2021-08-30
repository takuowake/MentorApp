package user

import (
	"github.com/gin-gonic/gin"
)

func UserEditDisplayAction(c *gin.Context){
	id := c.Param("id")    // URLからIDを取得して表示を分ける
	c.HTML(200, "user-edit.html", gin.H{
		"id": id,            // htmlファイルに値を渡す
	})
}