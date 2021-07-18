project-install:
	go get
	make docker-up
	make migrate
	make server-up


migrate:
	go run migrations/*.go migrate

docker-up:
	docker-compose up -d

server-up:
	make docker-up
	go run ./server.go

gqlgen-generate:
	go generate ./...