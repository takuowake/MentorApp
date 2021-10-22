package model

type Chatroom struct {
	ID uint64 `json:"id" gorm:"primaryKey; autoIncrement"`
	Username string
	PlanID uint64
	Plan Plan
	UserID uint64
	User User
	ReplyStatus string
}

type ChatroomUser struct {
	ID uint64 `json:"id" gorm:"primaryKey; autoIncrement"`
	ChatroomID uint64
	UserID uint64
	User User
}
