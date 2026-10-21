.PHONY: build build-cli run run-cli test test-race
build:
	@go build -o ./bin/DistributedCAS.elf ./cmd/

run: build
	@./bin/DistributedCAS.elf
	
test:
	@go test -v ./...

test-race:
	@go test -race ./...