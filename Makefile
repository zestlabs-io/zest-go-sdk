.PHONY: generate
generate:
	@./gen/generate.sh nocopy
	go get -u -f ./api/...

.PHONY: lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	@golangci-lint run ./...