package repository

import "database/sql"

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(DB *sql.DB) *CommentRepository {
	return &CommentRepository{
		db: DB,
	}
}
