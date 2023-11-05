package repository

import "database/sql"

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(DB *sql.DB) *PostRepository {
	return &PostRepository{
		db: DB,
	}
}
