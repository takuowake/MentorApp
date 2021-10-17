package server

import (
	"MentorApp/controller/admin"
	"MentorApp/controller/top"
	"MentorApp/controller/user"
	"MentorApp/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
	router := gin.Default()
	router.LoadHTMLGlob("views/static/*.html")

	router.GET("/", top.IndexDisplayAction)

	/*SessionCookieの設定*/
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/logout", model.LogoutAction)
	/*認証済みのアクセス可能なグループ*/
	authUserGroup := router.Group("/private")
	authUserGroup.Use(model.AuthRequired)
	{
		authUserGroup.GET("/me", me)
		authUserGroup.GET("/status", status)
	}

	router.GET("/login", model.LoginModel)
	router.POST("/login", model.LoginAction)
	router.GET("/signup", model.SignUpModel)
	router.POST("/signup", model.SignUpAction)
	//router.GET("/plan", plan.ListDisplayAction)
	router.GET("/user", user.UserPlanManageAction)
	router.GET("/user/add/:id", model.PlanModel)
	//router.POST("/user/edit/:id", model.CreatePlan)
	router.GET("/user/mypage", user.UserMyPageAction)
	router.GET("/user/message", user.UserMessageAction)
	router.GET("/user/payment", user.PaymentManageAction)
	router.GET("/user/resign", user.UserResignAction)
	router.GET("/admin", admin.AdminPlanManageAction)
	router.GET("/admin/mypage", admin.AdminMyPageAction)

	return router
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	c.JSON(http.StatusOK, gin.H{"user": username})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in user"})
}
