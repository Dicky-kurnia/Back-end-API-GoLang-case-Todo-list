package entity

import "time"

type Todo struct {
	ID              int       `gorm:"primaryKey" json:"id"`
	Title           string    `json:"title"`
	ActivityGroupID int       `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	UpdatedAt       time.Time `json:"updatedAt"`
	CreatedAt       time.Time `json:"createdAt"`
}
