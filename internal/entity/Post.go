package entity

import "time"

type Post struct {
	UserId    int       `json:"userId"`
	PostId    int       `json:"postId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
