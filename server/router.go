package server

import (
	"MentorApp/controller"
	"MentorApp/controller/admin"
	"MentorApp/controller/top"
	"MentorApp/controller/user"
	"MentorApp/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	/*Cors settings*/
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		AllowMethods: []string{
			"GET",
		},
		AllowHeaders: []string{
			"content-type",
		},
		AllowCredentials: true,
		MaxAge: 24 * time.Hour,
	}))
	router.LoadHTMLGlob("views/static/*.html")

	router.GET("/", top.IndexDisplayAction)

	/*SessionCookie settings*/
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/logout", controller.LogoutAction)
	/*認証済みのアクセス可能なグループ*/
	authUserGroup := router.Group("/private")
	authUserGroup.Use(model.AuthRequired)
	{
		authUserGroup.GET("/me", me)
		authUserGroup.GET("/status", status)
	}
	router.GET("/welcome", controller.ShowUser)
	router.GET("/login", model.LoginModel)
	router.POST("/login", controller.LoginAction)
	router.GET("/signup", model.SignUpModel)
	router.POST("/signup", controller.SignUpAction)
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
	router.GET("/api/v1/hello", func(c *gin.Context) {
		c.String(200, `{"message":"hello, hello, hello"}`)
	})

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
