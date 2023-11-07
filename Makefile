build:
	@swag init -d ./cmd
	@go build -o ./bin/sociot ./cmd/main.go

run:build
	@./bin/sociot