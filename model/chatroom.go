package model

type Chatroom struct {
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	UserID uint16
	User User
	ReplyStatus string
}

type ChatroomUser struct {
	RoomID uint16
	UserID uint16
	User User
}
