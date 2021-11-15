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

func ShowUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "hello"})
}


//ユーザー登録処理
func createUser(username string, password string) []error {
	passwordEncrypt, _ := model.PasswordEncrypt(password)
	/*Insert処理*/
	if err := database.DB.Create(model.User{Username: username, Password: passwordEncrypt}).GetErrors(); err != nil {
		return err
	}
	return nil
}

//ユーザーを一件取得
func getUser(username string) model.User {
	var user model.User
	database.DB.First(&user, "username = ?", username)
	return user
}

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
		if err := createUser(username, password); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		}
		c.Redirect(302, "/")
	}
}


//ユーザーログイン
func LoginAction(c * gin.Context) {
	//DBから取得したユーザーパスワード（Hash）
	dbPassword := getUser(c.PostForm("username")).Password
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
