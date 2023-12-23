package repository

import (
	"database/sql"
	"strings"
)

type VotesRepo struct {
	db *sql.DB
}

func NewVotesRepo(DB *sql.DB) VotesRepo {
	return VotesRepo{
		db: DB,
	}
}

func (repo VotesRepo) UpdatePostVotesById(postId int, userId int, voteType string) error {
	updateVoteValue := 0
	if strings.Compare(voteType, "u") == 0 {
		updateVoteValue = 1
	} else if strings.Compare(voteType, "d") == 0 {
		updateVoteValue = -1
	}

	var voteValue int
	query := `SELECT vote_type FROM votes WHERE postId = ? and userId = ?`
	err := repo.db.QueryRow(query, postId, userId).Scan(&voteValue)

	if err == sql.ErrNoRows {
		var postPresent int
		err = repo.db.QueryRow(`SELECT postId FROM posts WHERE postId = ?`, postId).Scan(&postPresent)
		if err == sql.ErrNoRows {
			return nil
		}
		query = `INSERT INTO votes (userId, postId, vote_type) VALUES (?, ?, ?)`
		_, err = repo.db.Exec(query, userId, postId, updateVoteValue)
		return err
	} else if err != nil {
		return err
	}

	if updateVoteValue == voteValue {
		updateVoteValue = 0
	}

	query = `UPDATE votes SET vote_type = ? WHERE userId = ? and postId = ?`
	_, err = repo.db.Exec(query, updateVoteValue, userId, postId)
	return err
}

func (repo VotesRepo) GetVotesCountById(postId int) (int, error) {
	query := `SELECT COALESCE(SUM(vote_type),0) AS count FROM votes WHERE postId = ?`
	var count int
	err := repo.db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo VotesRepo) GetUserVoted(postId int, userId int) (int, error) {
	query := `SELECT vote_type FROM votes WHERE postId = ? and userId = ?`
	var status int
	err := repo.db.QueryRow(query, postId, userId).Scan(&status)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return status, nil
}

func (repo VotesRepo) DeletePostVotesByPostId(postId int) {
	repo.db.Exec(`DELETE FROM votes WHERE postId = ?`, postId)
}

func (repo VotesRepo) DeletePostVotesByUserId(userId int) {
	repo.db.Exec(`DELETE FROM votes WHERE userId = ?`, userId)
}
