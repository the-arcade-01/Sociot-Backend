package entity

import (
	"database/sql"
	"time"
)

type Post struct {
	UserId    int       `json:"userId"`
	PostId    int       `json:"postId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostRequestBody struct {
	UserId  int    `json:"userId" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdatePostRequestBody struct {
	UserId  int    `json:"userId" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type DeletePostRequestBody struct {
	UserId int `json:"userId" validate:"required"`
}

func ScanIntoPost(rows *sql.Rows) (*Post, error) {
	post := new(Post)
	err := rows.Scan(
		&post.UserId,
		&post.PostId,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}
