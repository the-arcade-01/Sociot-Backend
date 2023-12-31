package repository

import (
	"database/sql"
	"fmt"
	"sociot/internal/entity"
	"sociot/internal/utils"
)

type PostRepository struct {
	db        *sql.DB
	votesRepo VotesRepo
}

func NewPostRepository(DB *sql.DB, repo VotesRepo) PostRepository {
	return PostRepository{
		db:        DB,
		votesRepo: repo,
	}
}

func (repo *PostRepository) GetPosts(sort string, tag string) ([]*entity.Post, error) {
	sortParam := "v.views"
	switch sort {
	case "new":
		sortParam = "p.createdAt"
	case "vote":
		sortParam = "votes"
	default:
		sortParam = "v.views"
	}
	var rows *sql.Rows
	if tag != "" {
		filter := "%" + tag + "%"
		query := fmt.Sprintf(utils.GET_POSTS_QUERY_TAGS, filter, sortParam)
		row, err := repo.db.Query(query)
		if err != nil {
			return nil, err
		}
		rows = row
	} else {
		query := fmt.Sprintf(utils.GET_POSTS_QUERY, sortParam)
		row, err := repo.db.Query(query)
		if err != nil {
			return nil, err
		}
		rows = row
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
	query := fmt.Sprintf(utils.GET_POST_BY_ID, postId)
	rows, err := repo.db.Query(query)
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
	result, err := repo.db.Exec(
		`INSERT INTO posts (userId, title, content) VALUES (?, ?, ?)`,
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

	err = repo.InsertPostView(postId)
	if err != nil {
		repo.DeletePostById(int(postId))
		return err
	}

	repo.InsertTags(postId, post.Tags)

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
	repo.DeleteTagsByPostId(post.PostId)
	repo.InsertTags(int64(post.PostId), post.Tags)

	return nil
}

func (repo *PostRepository) DeletePostById(postId int) error {
	repo.DeleteTagsByPostId(postId)

	err := repo.DeletePostViews([]int{postId})
	if err != nil {
		return err
	}

	repo.votesRepo.DeletePostVotesByPostId(postId)

	_, err = repo.db.Exec(
		`DELETE FROM posts WHERE postId = ?`,
		postId,
	)
	if err != nil {
		repo.InsertPostView(int64(postId))
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
	query := fmt.Sprintf(utils.GET_USERS_POSTS, userId)
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

func (repo *PostRepository) DeletePostViews(postIds []int) error {
	if len(postIds) == 0 {
		return nil
	}

	for _, postId := range postIds {
		_, err := repo.db.Exec(`DELETE FROM post_views WHERE postId = ?`, postId)
		if err != nil {
			return err
		}
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

	var postIds []int
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

	for _, postId := range postIds {
		repo.DeleteTagsByPostId(postId)
		repo.votesRepo.DeletePostVotesByPostId(postId)
	}

	repo.votesRepo.DeletePostVotesByUserId(userId)

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

func (repo *PostRepository) InsertPostView(postId int64) error {
	_, err := repo.db.Exec(
		`INSERT INTO post_views (postId) VALUE (?)`,
		postId,
	)
	return err
}

func (repo *PostRepository) InsertTags(postId int64, tags []string) {
	for _, tag := range tags {
		var tagId int64
		err := repo.db.QueryRow(`SELECT tagId FROM tags WHERE tag = ?`, tag).Scan(&tagId)
		if err != nil {
			result, err := repo.db.Exec(`INSERT INTO tags (tag) VALUE (?)`, tag)
			if err != nil {
				continue
			}
			tagId, err = result.LastInsertId()
			if err != nil {
				continue
			}
		}
		_, err = repo.db.Exec(
			`INSERT INTO post_tags (postId, tagId) VALUES (?, ?)`,
			postId, tagId)
		if err != nil {
			continue
		}
	}
}

func (repo *PostRepository) DeleteTagsByPostId(postId int) error {
	_, err := repo.db.Exec(`DELETE FROM post_tags WHERE postId = ?`, postId)
	return err
}
