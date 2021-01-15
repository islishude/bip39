.PHONY: update-wordlist
update-wordlist:
	@mkdir -p internal/wordlist
	@go run ./update-wordlist/main.go
	@goimports -w .

.PHONY: unit-test
unit-test:
	@go test -cover .
