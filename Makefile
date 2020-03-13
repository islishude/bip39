update-worldlist:
	go run tools/main.go
	goimports -w .
test:
	go test .
