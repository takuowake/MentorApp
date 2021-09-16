package model

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
)

//USERモデル
type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DBMS := os.Getenv("MentorApp_DBMS")
	USER := os.Getenv("MentorApp_USER")
	PASS := os.Getenv("MentorApp_PASS")
	DBNAME := os.Getenv("MentorApp_DBNAME")
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//DBの初期化
func dbInit() {
	db := gormConnect()
	/*Connection開放*/
	defer db.Close()
	db.AutoMigrate(&User{})
}

//ユーザー登録処理
func createUser(username string, password string) []error {
	dbInit()
	passwordEncrypt, _ := PasswordEncrypt(password)
	db := gormConnect()
	defer db.Close()
	/*Insert処理*/
	if err := db.Create(&User{Username: username, Password: passwordEncrypt}).GetErrors(); err != nil {
		return err
	}
	return nil
}

//ユーザーを一件取得
func getUser(username string) User {
	db := gormConnect()
	var user User
	db.First(&user, "username = ?", username)
	db.Close()
	return user
}

//ユーザー登録画面
func SignUpModel(c *gin.Context) {
	c.HTML(200, "signup.html", gin.H{})
}

//ユーザー登録
func SignUpAction(c *gin.Context) {
	var form User
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

//ユーザーログイン画面
func LoginModel(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

//ユーザーログイン
func LoginAction(c * gin.Context) {
	//DBから取得したユーザーパスワード（Hash）
	dbPassword := getUser(c.PostForm("username")).Password
	log.Println(dbPassword)
	//フォームから取得したユーザーパスワード
	formPassword := c.PostForm("password")

	//ユーザーパスワードの比較
	if err := CompareHashAndPassword(dbPassword, formPassword); err != nil {
		log.Println("ログインできませんでした")
	    c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
		c.Abort()
	} else {
		log.Println("ログインできました")
		c.Redirect(302, "/")
	}
}
