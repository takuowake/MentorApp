package server

import (
	"MentorApp/controllers"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	router.GET("/", controllers.IndexDisplayAction)
	router.GET("/user", controllers.UserMessageDisplayAction)
	router.GET("/user/plan/add", controllers.UserAddDisplayAction)
	router.GET("/user/plan/edit/:id", controllers.UserEditDisplayAction)

	return router
}