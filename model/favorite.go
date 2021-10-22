package model

type Favorite struct {
	ID uint64 `json:"id" gorm:"primaryKey; autoIncrement"`
	UserID uint64
	User User
	PlanID uint64
	Plan Plan
}
