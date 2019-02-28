GO := go
CMD_PATH := $(shell go list -m)/cmd/barr
BUILD_DIR := ./build/package
BIN := $(BUILD_DIR)/barr
VERSION := $(shell git describe --tags)

all:
	@$(GO) build -ldflags '-X $(CMD_PATH).Version=$(VERSION) -s -w' -o $(BIN) $(CMD_PATH)
	@upx $(BIN)
	@ln -sf $(PWD)/init/barr.service $(BUILD_DIR)/
	@cd $(BUILD_DIR); holo-build --force --format=pacman holo.toml

clean:
	rm -f $(BIN) $(wildcard $(BUILD_DIR)/*.pkg.tar.xz)
