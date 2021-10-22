package model

import "time"

type Review struct {
	UserID uint64
	User User
	PlanID uint64
	Plan Plan
	IsPublish time.Time
	Content string `form:"content" binding:"required"`
	Rank uint8
}
