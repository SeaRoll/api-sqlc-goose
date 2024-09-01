gen:
	go generate ./...

build:
	go build -o api-server cmd/main/main.go
