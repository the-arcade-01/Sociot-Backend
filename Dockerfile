FROM golang:1.21.3

WORKDIR /usr/src/sociot

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag
RUN swag init -g ./cmd/main.go -o ./docs --parseInternal true
RUN go build -o ./bin/sociot ./cmd/main.go

RUN chmod +x ./bin/sociot
RUN chmod +x ./scripts/run.sh

ENTRYPOINT [ "sh", "./scripts/run.sh" ]