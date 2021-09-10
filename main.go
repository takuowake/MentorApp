package main

import (
	"MentorApp/server"
	"github.com/gin-gonic/gin"
)

func main() {
	router := server.GetRouter()

	//ログイン
	authorized := router.Group("/user", gin.BasicAuth(gin.Accounts{
		"username":    "password",
		"username2":   "password2",
	}))
	authorized.GET("/login", func(c *gin.Context) {
		//user := c.MustGet(gin.AuthUserKey).(string)
		//c.JSON(200, gin.H{ "message": "Hello " + user })
		c.HTML(200, "index.html", gin.H{})
	})

	router.Run(":8080")
}
