package entity

import (
	"database/sql"
	"log"
	"time"
)

type User struct {
	UserId    int       `json:"userId"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserDetails struct {
	UserId    int       `json:"userId"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserSearch struct {
	UserId    int       `json:"userId"`
	UserName  string    `json:"username"`
	PostCount int       `json:"postCount"`
	ViewCount int       `json:"viewCount"`
	Votes     int       `json:"votes"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserLoginDetails struct {
	Token  string `json:"token"`
	UserId int    `json:"userId"`
}

type UpdateUserNameReqBody struct {
	UserName string `json:"username" validate:"required,min=4"`
}

type UpdateUserPasswordReqBody struct {
	Password string `json:"password" validate:"required,min=4"`
}

type CreateUserRequestBody struct {
	UserName string `json:"username" validate:"required,min=4"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type LoginUserRequestBody struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ScanIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.UserId,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Printf("error occurred while scanning db row into user, %v\n", err)
		return nil, err
	}
	return user, nil
}

func ScanIntoUserSearch(rows *sql.Rows) (*UserSearch, error) {
	user := new(UserSearch)
	err := rows.Scan(
		&user.UserId,
		&user.UserName,
		&user.PostCount,
		&user.ViewCount,
		&user.Votes,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
