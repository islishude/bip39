.PHONY: update-wordlist
update-wordlist:
	@mkdir -p wordlist
	@go run github.com/islishude/bip39/update-wordlist
	@goimports -w .

.PHONY: unit-test
unit-test:
	@go test -cover .
