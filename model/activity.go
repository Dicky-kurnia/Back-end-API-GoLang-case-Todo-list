package model

import "time"

type ActivityInputID struct {
	ID int `uri:"id" binding:"required"`
}

type ActivityInput struct {
	Title      string    `json:"title" binding:"required"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"craetedAt"`
	Updated_At time.Time `json:"updatedAt"`
}
