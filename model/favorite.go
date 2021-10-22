package model

type Favorite struct {
	ID uint64 `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	UserID uint64
	User User
	PlanID uint64
	Plan Plan
}
