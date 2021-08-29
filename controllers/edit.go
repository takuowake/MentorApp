package controllers

import "github.com/gin-gonic/gin"

func UserEditDisplayAction(c *gin.Context){
	id := c.Param("id")    //URLからIDを取得して表示を分ける
	c.HTML(200, "user-plan-edit.html", gin.H{
		"id": id,               //htmlファイルにあたいを渡す
	})
}
