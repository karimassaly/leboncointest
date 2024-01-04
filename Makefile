run:
	@go run ./cmd/api/main.go

store:
	@docker-compose up --build

