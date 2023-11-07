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

func (repo *UserRepository) GetUsers() []*entity.User {
	user := &entity.User{
		UserId:   1,
		UserName: "meowth",
		Email:    "meowth@gmail.com",
	}
	var users []*entity.User
	users = append(users, user)
	return users
}

func (repo *UserRepository) GetUserById(userId int) *entity.User {
	user := &entity.User{
		UserId:   userId,
		UserName: "ben10",
		Email:    "ben10@gmail.com",
	}
	return user
}
