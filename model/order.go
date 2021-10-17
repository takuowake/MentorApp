package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID string `gorm:"primaryKey"`
	UserID   uint16
	User     User
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrdersDetail struct {
	ID string `gorm:"primaryKey"`
	PlanID string
	Plan Plan
	OrderID string
	Order Order
	Quantity uint8
}
