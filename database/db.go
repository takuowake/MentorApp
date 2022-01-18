package database

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DBMS := os.Getenv("MENTOR_APP_DBMS")
	USER := os.Getenv("MENTOR_APP_USER")
	PASS := os.Getenv("MENTOR_APP_PASS")
	DBNAME := os.Getenv("MENTOR_APP_DBNAME")
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// DbInit initialize Db.
func DbInit() *gorm.DB {
	db := gormConnect()
	db.LogMode(true)
	DB = db
	return db
}

// GetDB gets a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
