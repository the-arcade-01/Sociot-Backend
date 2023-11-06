package entity

type User struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
