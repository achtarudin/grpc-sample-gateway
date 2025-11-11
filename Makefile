BIN_FILENAME=grpc-sample-gateway

.PHONY: install-tools
install-tools:
	go install github.com/mitranim/gow@latest

.PHONY: install-deps
install-deps:
	go mod tidy
	go mod download
	go get -u github.com/achtarudin/grpc-sample

.PHONY: clean-server-bin
clean-server-bin:
	rm -rf ./bin

.PHONY: build-server
build-server: clean-server-bin install-deps
	go build -o ./bin/$(BIN_FILENAME) ./cmd/server

.PHONY: prod-server
prod-server: build-server
	./bin/$(BIN_FILENAME)

.PHONY: dev-server
dev-server: 
	gow run ./cmd/server
