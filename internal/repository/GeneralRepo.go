package repository

import (
	"database/sql"
	"fmt"
	"sociot/internal/entity"
	"sociot/internal/utils"
)

type GeneralRepository struct {
	db *sql.DB
}

func NewGeneralRepository(db *sql.DB) GeneralRepository {
	return GeneralRepository{
		db: db,
	}
}

func (repo *GeneralRepository) Search(option string, search string) (*entity.SearchResults, error) {
	var users []*entity.UserSearch
	var posts []*entity.Post

	search = "%" + search + "%"
	if option == "user" {
		query := fmt.Sprintf(utils.SEARCH_USER, search)
		rows, err := repo.db.Query(query)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			user, err := entity.ScanIntoUserSearch(rows)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}
	} else {
		query := fmt.Sprintf(utils.SEARCH_POST, search)
		rows, err := repo.db.Query(query)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			post, err := entity.ScanIntoPost(rows)
			if err != nil {
				return nil, err
			}
			posts = append(posts, post)
		}
	}

	results := entity.NewSearchResults(posts, users)
	return results, nil
}
