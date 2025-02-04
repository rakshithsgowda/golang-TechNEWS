tidy:
	go mod tidy

run:
	@go run ./cmd/web/


run/migrate:
	@go run ./cmd/web -migrate=true