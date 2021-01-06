build:
	go build -o api

debug:
	go run main.go 

test:
	go test ./...