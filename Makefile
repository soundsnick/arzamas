build:
	go build -o api

dev:
	go run main.go 

test:
	go test ./...