.PHONY: update-wordlist
update-wordlist:
	@mkdir -p internal/wordlist
	@go run ./update-wordlist/main.go
	@gofmt -w internal/wordlist

.PHONY: unit-test
unit-test:
	@go test -cover .

pre-commit:
	@go build ./...
	@go vet ./...
	@go mod tidy
	@go fmt ./...
	@go fix ./...
	@go test -race -cover .
