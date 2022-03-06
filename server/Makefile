.PHONY: server clean-server

SERVER_DIR := server
SERVER_BIN_DIR := bin
SERVER_MAIN_FILE := $(SERVER_DIR)/$(SERVER_BIN_DIR)/main.go
SERVER_TARGET_FILE := $(SERVER_DIR)
SERVER_TARGET := $(OUT_DIR)/$(SERVER_TARGET_FILE)

$(SERVER_TARGET): FORCE
	go build -o $(SERVER_TARGET) $(SERVER_MAIN_FILE)

server: $(SERVER_TARGET)

clean-server:
	-rm $(SERVER_TARGET)

clean: clean-server
