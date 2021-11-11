include .env

run:
	HOST=$(HOST) PORT=$(PORT) USER=$(USER) PASSWORD=$(PASSWORD) DB_NAME=$(DB_NAME) go run main.go
