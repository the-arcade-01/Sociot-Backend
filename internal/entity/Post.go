package entity

import (
	"database/sql"
	"strings"
	"time"
)

type Post struct {
	UserId    int       `json:"userId"`
	UserName  string    `json:"username"`
	PostId    int       `json:"postId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	Views     int       `json:"views"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostRequestBody struct {
	UserId  int    `json:"userId" validate:"required"`
	Title   string `json:"title" validate:"required,min=4"`
	Content string `json:"content" validate:"required,min=4"`
}

type UpdatePostRequestBody struct {
	UserId  int    `json:"userId" validate:"required"`
	Title   string `json:"title" validate:"required,min=4"`
	Content string `json:"content" validate:"required,min=4"`
}

type DeletePostRequestBody struct {
	UserId int `json:"userId" validate:"required"`
}

func ScanIntoPost(rows *sql.Rows) (*Post, error) {
	post := new(Post)
	var tagsList string

	err := rows.Scan(
		&post.UserId,
		&post.UserName,
		&post.PostId,
		&post.Title,
		&post.Content,
		&tagsList,
		&post.Views,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	tags := strings.Split(tagsList, ",")
	post.Tags = tags

	return post, nil
}

func ScanIntoTag(rows *sql.Rows) (string, error) {
	var tag string
	err := rows.Scan(
		&tag,
	)
	if err != nil {
		return "", err
	}
	return tag, nil
}
