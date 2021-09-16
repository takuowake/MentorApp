package server

import (
	"MentorApp/controller/admin"
	"MentorApp/controller/plan"
	"MentorApp/controller/top"
	"MentorApp/controller/user"
	"MentorApp/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"net/http"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")


	router.GET("/", top.IndexDisplayAction)
	/*SessionCookieの設定*/
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	/*認証済みのアクセス可能なグループ*/
	authUserGroup := router.Group("/auth")
	authUserGroup.Use(model.AuthRequired())
	{
		authUserGroup.GET("/getSample", handler.getSample)
	}

	router.GET("/login", model.LoginModel)
	router.POST("/login", model.LoginAction)
	router.GET("/signup", model.SignUpModel)
	router.POST("/signup", model.SignUpAction)
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

func me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in user"})
}
