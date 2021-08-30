package user

import (
"github.com/gin-gonic/gin"
)

func UserAddDisplayAction(c *gin.Context){
	id := c.Param("id")    // URLからIDを取得して表示を分ける
	c.HTML(200, "user-add.html", gin.H{
		"id": id,            // htmlファイルに値を渡す
	})
}