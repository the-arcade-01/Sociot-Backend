package repository

import (
	"database/sql"
	"sociot/internal/entity"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(DB *sql.DB) CommentRepository {
	return CommentRepository{
		db: DB,
	}
}

func (repo *CommentRepository) GetCommentById() entity.Comment {
	comment := entity.Comment{
		UserId:  1,
		PostId:  1,
		Content: "This is my first comment!!",
	}
	return comment
}
