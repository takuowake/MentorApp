package main

import (
	"MentorApp/database"
	"MentorApp/server"
	"github.com/jinzhu/gorm"
)

func main() {
	db := database.DbInit()
	Migrate(db)
	defer db.Close()
	router := server.GetRouter()

	router.Run(":8090")
}

func Migrate(db *gorm.DB) {
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Category{})
	//db.AutoMigrate(&model.OrdersDetail{})
	//db.AutoMigrate(&model.Plan{})
	//db.AutoMigrate(&model.Order{})
	//db.AutoMigrate(&model.Review{})
	//db.AutoMigrate(&model.Favorite{})
	//db.AutoMigrate(&model.Chatroom{})
	//db.AutoMigrate(&model.ChatroomUser{})
}
