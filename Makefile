server-run:
	docker-compose up -d
	go run ./server.go

gqlgen-generate:
	go generate ./...