CMD_PATH := $(shell GO111MODULE=on go list -m)/cmd/barr
BUILD_DIR := $(abspath ./out)
BIN := $(BUILD_DIR)/barr
VERSION := $(shell git describe --tags)

all: $(BIN)

$(BIN):
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags '-X main.Version=$(VERSION) -s -w' -o $(BIN) $(CMD_PATH)
	@upx $(BIN)

clean:
	rm -rf $(BUILD_DIR)

ifdef GOBIN
INSTALL_DIR := $(GOBIN)

install: $(BIN)
	mv $(BIN) $(INSTALL_DIR)

uninstall:
	rm -f $(INSTALL_DIR)/barr
endif



.PHONY: all install uninstall clean
