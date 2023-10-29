build:
	@go build -o ./bin/sociot ./cmd/main.go

run:build
	@./bin/sociot