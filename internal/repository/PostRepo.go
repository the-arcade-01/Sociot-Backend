package repository

import (
	"database/sql"
	"sociot/internal/entity"
	"strings"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(DB *sql.DB) PostRepository {
	return PostRepository{
		db: DB,
	}
}

func (repo *PostRepository) GetPosts() ([]*entity.Post, error) {
	query := `
	SELECT 
    u.userId,
    u.userName,
    p.postId,
    p.title,
    p.content,
    v.views,
    p.createdAt,
    p.updatedAt
FROM 
    posts p
LEFT JOIN 
    post_views v ON p.postId = v.postId
LEFT JOIN 
    users u ON p.userId = u.userId
ORDER BY 
    views DESC; 
	`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	var posts []*entity.Post
	for rows.Next() {
		post, err := entity.ScanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo *PostRepository) GetPostById(postId int) (*entity.Post, error) {
	query := `
	SELECT 
    u.userId,
    u.userName,
    p.postId,
    p.title,
    p.content,
    v.views,
    p.createdAt,
    p.updatedAt
FROM 
    posts p
LEFT JOIN 
    post_views v ON p.postId = v.postId
LEFT JOIN 
    users u ON p.userId = u.userId
WHERE
    p.postId = ?
	`
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

func (repo *PostRepository) CreatePost(post *entity.Post) error {
	query := `INSERT INTO posts (userId, title, content) VALUES (?, ?, ?)`
	result, err := repo.db.Exec(
		query,
		post.UserId,
		post.Title,
		post.Content,
	)
	if err != nil {
		return err
	}

	postId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	query = `INSERT INTO post_views (postId) VALUE (?)`
	_, err = repo.db.Exec(
		query,
		postId,
	)
	if err != nil {
		repo.DeletePostById(int(postId))
		return err
	}

	return nil
}

func (repo *PostRepository) UpdatePostById(post *entity.Post) error {
	query := `UPDATE posts SET title = ?, content = ? WHERE postId = ?`
	_, err := repo.db.Exec(
		query,
		post.Title,
		post.Content,
		post.PostId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) DeletePostById(postId int) error {
	query := `DELETE FROM post_views WHERE postId = ?`
	_, err := repo.db.Exec(
		query,
		postId,
	)

	if err != nil {
		return err
	}

	query = `DELETE FROM posts WHERE postId = ?`
	_, err = repo.db.Exec(
		query,
		postId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) UpdatePostViewsById(postId int) error {
	query := `UPDATE post_views SET views = views + 1 WHERE postId = ?`
	_, err := repo.db.Exec(
		query,
		postId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) GetUserPosts(userId int) ([]*entity.Post, error) {
	query := `
	SELECT 
    u.userId,
    u.userName,
    p.postId,
    p.title,
    p.content,
    v.views,
    p.createdAt,
    p.updatedAt
FROM 
    posts p
LEFT JOIN 
    post_views v ON p.postId = v.postId
LEFT JOIN 
    users u ON p.userId = u.userId
WHERE
    u.userId = ?
	`
	rows, err := repo.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	var posts []*entity.Post
	for rows.Next() {
		post, err := entity.ScanIntoPost(rows)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo *PostRepository) DeletePostViews(postIds []interface{}) error {
	if len(postIds) == 0 {
		return nil
	}

	query := `DELETE FROM post_views WHERE postId IN (`
	query += strings.Repeat("?, ", len(postIds)-1) + "?)"

	_, err := repo.db.Exec(query, postIds...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) DeletePostByUserId(userId int) error {
	query := `SELECT postId FROM posts WHERE userId = ?`
	rows, err := repo.db.Query(query, userId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var postIds []interface{}
	for rows.Next() {
		var postId int
		err := rows.Scan(&postId)
		if err != nil {
			return err
		}
		postIds = append(postIds, postId)
	}

	err = repo.DeletePostViews(postIds)
	if err != nil {
		return err
	}

	query = `DELETE FROM posts WHERE userId = ?`
	_, err = repo.db.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PostRepository) GetTags() ([]string, error) {
	query := `
	SELECT
	t.tag
FROM
	tags t 
LEFT JOIN 
	post_tags pt ON pt.tagId = t.tagId
GROUP BY 
 	t.tagId 
ORDER BY 
	COUNT(pt.tagId) 
DESC
	`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	var tags []string
	for rows.Next() {
		tag, err := entity.ScanIntoTag(rows)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
