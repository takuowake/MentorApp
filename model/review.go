package model

import "time"

type Review struct {
	UserID uint16
	User User

	IsPublish time.Time
	Content string `form:"content" binding:"required"`
	Rank uint8
}
