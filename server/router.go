package server

import (
	"MentorApp/controller/admin"
	"MentorApp/controller/login"
	"MentorApp/controller/plan"
	"MentorApp/controller/signup"
	"MentorApp/controller/top"
	"MentorApp/controller/user"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	router.GET("/", top.IndexDisplayAction)
	router.GET("/login", login.LoginAction)
	router.GET("/signup", signup.SignUpAction)
	router.GET("/plan", plan.ListDisplayAction)
	router.GET("/user", user.UserPlanManageAction)
	router.GET("/user/add/:id", user.UserAddDisplayAction)
	router.GET("/user/edit/:id", user.UserEditDisplayAction)
	router.GET("/user/mypage", user.UserMyPageAction)
	router.GET("/user/message", user.UserMessageAction)
	router.GET("/user/payment", user.PaymentManageAction)
	router.GET("/user/resign", user.UserResignAction)
	router.GET("/admin", admin.AdminPlanManageAction)
	router.GET("/admin/mypage", admin.AdminMyPageAction)

	return router
}