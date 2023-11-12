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

func (repo *PostRepository) GetPostById(postId int) (*entity.Post, error) {
	query := `SELECT * FROM posts WHERE id = ?`
	rows, err := repo.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	var post *entity.Post
	for rows.Next() {
		post, err = entity.ScanIntoPost(rows)
		if err != nil {
			return nil, err
		}
	}
	return post, nil
}

func (repo *PostRepository) CreatePost(post *entity.PostRequestBody) error {
	query := `INSERT INTO posts (userId, content) VALUES (?, ?)`
	_, err := repo.db.Exec(
		query,
		post.UserId,
		post.Content,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) UpdatePostById(postId int, post *entity.UpdatePostRequestBody) error {
	query := `UPDATE posts SET content = ? WHERE id = ?`
	_, err := repo.db.Exec(
		query,
		post.Content,
		postId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) DeletePostById(postId int) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := repo.db.Exec(
		query,
		postId,
	)
	if err != nil {
		return err
	}
	return nil
}
