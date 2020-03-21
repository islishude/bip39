update-worldlist:
	@go run tools/main.go
	@goimports -w .

unit-test:
	go test -cover .
