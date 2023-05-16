package model

type TodoRequest struct {
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

type TodoResponse struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	ActivityGroupID int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
	UpdatedAt       string `json:"updatedAt"`
	CreatedAt       string `json:"createdAt"`
}
