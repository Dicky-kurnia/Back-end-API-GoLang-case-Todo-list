package entity

import "time"

type Activity struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
