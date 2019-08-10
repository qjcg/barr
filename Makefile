CMD_PATH := $(shell go list -m)/cmd/barr
BUILD_DIR := $(abspath ./build)
BIN := $(BUILD_DIR)/barr
VERSION := $(shell git describe --tags)

all: $(BIN)

$(BIN):
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags '-X $(CMD_PATH).Version=$(VERSION) -s -w' -o $(BIN) $(CMD_PATH)
	@upx $(BIN)

install: $(BIN)
	mv $(BIN) $(GOBIN)

uninstall:
	rm -f $(GOBIN)/barr

clean:
	rm -rf $(BUILD_DIR)
