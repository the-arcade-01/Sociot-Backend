package repository

import (
	"database/sql"
	"sociot/internal/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return UserRepository{
		db: DB,
	}
}

func (repo *UserRepository) GetUsers() ([]*entity.User, error) {
	query := `SELECT * FROM users`
	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}

	var users []*entity.User
	for rows.Next() {
		user, err := entity.ScanIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *UserRepository) GetUserById(userId int) *entity.User {
	user := &entity.User{
		UserId:   userId,
		UserName: "ben10",
		Email:    "ben10@gmail.com",
	}
	return user
}

func (repo *UserRepository) UpdateUserById(userId int, userBody *entity.UpdateUserRequestBody) error {
	return nil
}

func (repo *UserRepository) DeleteUserById(userId int) error {
	return nil
}

func (repo *UserRepository) CreateUser(userBody *entity.CreateUserRequestBody) error {
	return nil
}

func (repo *UserRepository) LoginUser(userBody *entity.LoginUserRequestBody) error {
	return nil
}
