## Sociot Backend

Build in Go

## Installation

```
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
go install github.com/swaggo/swag/cmd/swag
go get github.com/swaggo/http-swagger/v2
go get github.com/go-chi/jwtauth/v5
go get github.com/joho/godotenv
go get github.com/go-sql-driver/mysql
go get golang.org/x/crypto/bcrypt
go get -u github.com/go-playground/validator/v10
```

## Project Structure

```shell
.
├── bin
│   └── sociot
├── cmd
│   └── main.go
├── config
│   ├── config.go
│   ├── DBConfig.go
│   └── JWTConfig.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── controller
│   │   ├── CommentController.go
│   │   ├── GreetController.go
│   │   ├── PostController.go
│   │   └── UserController.go
│   ├── entity
│   │   ├── Comment.go
│   │   ├── Post.go
│   │   ├── Response.go
│   │   └── User.go
│   ├── repository
│   │   ├── CommentRepo.go
│   │   ├── PostRepo.go
│   │   └── UserRepo.go
│   ├── service
│   │   ├── CommentService.go
│   │   ├── PostService.go
│   │   └── UserService.go
│   └── utils
│       ├── CommonUtils.go
│       └── Constants.go
├── Makefile
├── notes
├── Readme.md
└── scripts
    ├── db.sql
    └── setup.sh
```

## API Design

```js
POST /v1/user/login
POST /v1/user/
GET /v1/user/{id}
PUT /v1/user/{id}
DELETE /v1/user/{id}
GET /v1/posts/users/{id}
GET /v1/comments/users/{id}
GET /v1/posts/tags
GET /v1/posts?sortBy=views&tag=tech
POST /v1/posts
GET /v1/posts/{id}
PUT /v1/posts/{id}
DELETE /v1/posts/{id}
GET /v1/posts/{postId}/comments?sortBy=votes
POST /v1/posts/{postId}/comments
GET /v1/comments/{commentId}
PUT /v1/comments/{commentId}
DELETE /v1/comments/{commentId}
```
