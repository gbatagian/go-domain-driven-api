run:
	source .env && go run .

test:
	go test ./...

coverage:
	go test -cover ./...

restart:
	docker compose down && docker network prune && docker compose up