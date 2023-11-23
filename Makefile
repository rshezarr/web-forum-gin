.SILENT:

build:
	go build -o forum ./cmd/app/main.go
	./forum

run-compose:
	docker-compose up -d --build

stop-compose:
	docker-compose down

run-psql:
	docker run --name=forum -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres:13-alpine