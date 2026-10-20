.PHONY: build build-cli run run-cli test test-race
build:
	@go build -o ./bin/DistributedCAS.elf ./cmd/

run: build
	@./bin/DistributedCAS
	
test:
	@go test -v ./...

test-race:
	@go test -race ./...