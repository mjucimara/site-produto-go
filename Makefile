APP_NAME=site-produto
CMD_DIR=./cmd/web
BIN_DIR=./bin

.PHONY: run check build clean

run:
	go run $(CMD_DIR)

check:
	go fmt ./...
	go test ./...

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

clean:
	rm -rf $(BIN_DIR)
