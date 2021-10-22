package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID uint64 `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	UserID   uint64
	User     User
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrdersDetail struct {
	ID uint64 `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	PlanID string
	Plan Plan
	OrderID string
	Order Order
	Quantity uint8
}
