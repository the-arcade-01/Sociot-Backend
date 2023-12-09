## Sociot Backend

Build in Go

## Installation

```bash
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

## Docker cmds

```bash
docker compose build
docker compose up
```

If you face some corrupt db issue

```bash
docker image rm --force <app-container-id> <mysql-container-id>

# remove the container to which the volume is attached
docker rm -v <container-id>

docker volume rm --force <volume-name>
```

If you face db migrations issue

```bash
docker exec -it <mysql-container-id> sh

# cd into below folder and check whether sql file is present or not
cd docker-entrypoint-initdb.d/

# run below cmd to populate tables
mysql -u <user> -p < ./<sql-file>
```
