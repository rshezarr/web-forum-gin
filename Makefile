.SILENT:

build: 
	go build -o forum ./cmd/app/main.go
	./forum

run: 
	docker compose up --build -d

dbrun:
	docker run --name=forum -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres