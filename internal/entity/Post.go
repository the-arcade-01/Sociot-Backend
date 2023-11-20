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
	UserId  int    `json:"userId"`
	Content string `json:"content"`
}

type UpdatePostRequestBody struct {
	UserId  int    `json:"userId"`
	Content string `json:"content"`
}

type DeletePostRequestBody struct {
	UserId int `json:"userId"`
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
