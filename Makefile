run:
	source .env && go run .

test:
	go test ./...

coverage:
	go test -cover ./...