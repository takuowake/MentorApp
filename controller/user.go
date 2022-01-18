package controller

import (
	"MentorApp/database"
	"MentorApp/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

//ユーザー登録
func SignUpAction(c *gin.Context) {
	var form model.User
	//バリデーション処理
	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		c.Abort()
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		//登録ユーザーが重複していた場合にはじく処理
		if err := model.CreateUser(username, password); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		}
		c.Redirect(302, "/")
	}
}


//ユーザーログイン
func LoginAction(c *gin.Context) {
	//DBから取得したユーザーパスワード（Has
	//h）
	dbPassword := model.GetUser(c.PostForm("username")).Password
	log.Println(dbPassword)
	//フォームから取得したユーザーパスワード
	formPassword := c.PostForm("password")
	//ユーザーパスワードの比較
	if err := model.CompareHashAndPassword(dbPassword, formPassword); err != nil {
		log.Println("ログインできませんでした")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to authenticated"})
		return
		c.Abort()
	} else {
		log.Println("ログインできました")
		c.Redirect(302, "/")

	}

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate from input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Save the username in the session
	session.Set("user", username) // In real world usage you'd set this to users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func SimpleLoginAction(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"message": "s"})
	var login model.Login
	var user model.User
	c.BindJSON(&login)

	if login.Email != "" && login.Password != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		database.DB.Where("username = ?", login.Email).First(&user)		// Display error
		if (model.User{} == user) {
			c.JSON(401, gin.H{"error": "Email Address is not found."})
		}
		c.JSON(200, gin.H{"user": user})
	} else {
		// Display error
		c.JSON(401, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users

}





// LogoutAction is a handler that parses a form and checks for specific data
func LogoutAction(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete("user")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	log.Println("ログアウトできました")
}
