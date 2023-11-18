package repository

import (
	"database/sql"
	"errors"
	"sociot/internal/entity"
	"sociot/internal/utils"
	"strconv"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return UserRepository{
		db: DB,
	}
}

func (repo *UserRepository) GetUsers() ([]*entity.UserDetails, error) {
	query := `SELECT * FROM users`
	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}

	var users []*entity.UserDetails
	for rows.Next() {
		user, err := entity.ScanIntoUser(rows)
		userDetails := &entity.UserDetails{
			UserId:    user.UserId,
			UserName:  user.UserName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		if err != nil {
			return nil, err
		}
		users = append(users, userDetails)
	}
	return users, nil
}

func (repo *UserRepository) GetUserById(userId int) (*entity.UserDetails, error) {
	query := `SELECT * FROM users WHERE userId = ?`
	rows, err := repo.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	var user *entity.User
	for rows.Next() {
		user, err = entity.ScanIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		return nil, nil
	}

	userDetails := &entity.UserDetails{
		UserId:    user.UserId,
		Email:     user.Email,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userDetails, nil
}

func (repo *UserRepository) UpdateUserById(userId int, user *entity.User) error {
	hashPassword, err := utils.GetHashPassword(user.Password)
	if err != nil {
		return err
	}

	query := `UPDATE users SET userName = ?, password = ? WHERE userId = ?`
	_, err = repo.db.Exec(
		query,
		user.UserName,
		hashPassword,
		userId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUserById(userId int) error {
	query := `DELETE FROM users WHERE userId = ?`
	_, err := repo.db.Exec(
		query, userId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) CreateUser(user *entity.User) error {
	hashPassword, err := utils.GetHashPassword(user.Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (userName, email, password) VALUES (?, ?, ?)`
	_, err = repo.db.Exec(
		query,
		user.UserName,
		user.Email,
		hashPassword,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) LoginUser(user *entity.User) (*entity.UserDetails, error) {
	query := `SELECT * FROM users WHERE email = ?`
	rows, err := repo.db.Query(query, user.Email)
	if err != nil {
		return nil, err
	}

	var dbUser *entity.User
	for rows.Next() {
		dbUser, err = entity.ScanIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if dbUser == nil {
		return nil, errors.New("user doesn't exists, please create an account")
	}

	if err := utils.CheckPassword(dbUser.Password, user.Password); err != nil {
		return nil, errors.New("incorrect password, please try again")
	}

	userDetails := &entity.UserDetails{
		UserId:    dbUser.UserId,
		UserName:  dbUser.UserName,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}

	return userDetails, nil
}

func (repo *UserRepository) CheckExistingUser(user *entity.User) error {
	query := `SELECT COUNT(userId) FROM users WHERE userName = ?`
	var records string
	err := repo.db.QueryRow(query, user.UserName).Scan(&records)
	if err != nil {
		return errors.New("please use different username")
	}
	recordsNum, err := strconv.Atoi(records)

	if recordsNum != 0 || err != nil {
		return errors.New("username already taken, please use different username")
	}

	query = `SELECT COUNT(email) FROM users WHERE email = ?`
	err = repo.db.QueryRow(query, user.Email).Scan(&records)
	if err != nil {
		return errors.New("please use different email")
	}
	recordsNum, err = strconv.Atoi(records)

	if recordsNum != 0 || err != nil {
		return errors.New("email already exists, please use different email")
	}

	return nil
}
