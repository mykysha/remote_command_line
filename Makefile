go-formatter:
	gofumpt -l -s -w .

go-server:
	go run cmd/server/main.go

go-client:
	go run cmd/client/main.go
