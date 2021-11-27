package model

import (
	"MentorApp/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"time"
)

//USERモデル
type User struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey; autoIncrement"`
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
	EmailAddress string
	FirstName string
	FirstNameKana string
	LastName string
	LastNameKana string
	LoginAt time.Time
	Birthday time.Time
	Occupation string
	Sex string
	Rank uint8
}

//ユーザー登録画面
func SignUpModel(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

//ユーザーログイン画面
func LoginModel(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}


func ShowUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "user"})
}


//ユーザー登録処理
func CreateUser(username string, password string) []error {
	passwordEncrypt, _ := PasswordEncrypt(password)
	/*Insert処理*/
	if err := database.DB.Create(User{Username: username, Password: passwordEncrypt}).GetErrors(); err != nil {
		return err
	}
	return nil
}

//ユーザーを一件取得
func GetUser(username string) User {
	var user User
	database.DB.First(&user, "username = ?", username)
	return user
}
