## Sociot Backend

Build in Go

## Installation

```
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
```

## Project Structure

```shell
.
├── bin
│ └── sociot
├── cmd
│ └── main.go
├── config
│ └── config.go
├── docs
├── go.mod
├── internal
│ ├── controller
│ ├── entity
│ ├── repository
│ └── service
├── Makefile
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
GET /v1/tags
GET /v1/posts?sortBy=views&tag=tech
POST /v1/posts
GET /v1/posts/{id}
PUT /v1/posts/{id}
DELETE /v1/posts/{id}
GET /v1/posts/{postId}/comments?sortBy=votes
POST /v1/posts/{postId}/comments
PUT /v1/comments/{commentId}
DELETE /v1/comments/{commentId}
```
