package entity

type Post struct {
	UserId    int    `json:"userId"`
	PostId    int    `json:"postId"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

