package repository

import (
	"database/sql"
	"sociot/internal/entity"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(DB *sql.DB) PostRepository {
	return PostRepository{
		db: DB,
	}
}

func (repo *PostRepository) GetPosts() []*entity.Post {
	post := &entity.Post{
		UserId:  1,
		PostId:  1,
		Content: "this is my first post!!",
	}
	var posts []*entity.Post
	posts = append(posts, post)
	return posts
}
