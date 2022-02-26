.PHONY: queue clean-queue

OUT_DIR := out

QUEUE_DIR := queue
QUEUE_BIN_DIR := bin
QUEUE_MAIN_FILE := $(QUEUE_DIR)/$(QUEUE_BIN_DIR)/main.go
QUEUE_TARGET_FILE := $(QUEUE_DIR)
QUEUE_TARGET := $(OUT_DIR)/$(QUEUE_TARGET_FILE)

$(QUEUE_TARGET): FORCE
	go build -o $(QUEUE_TARGET) $(QUEUE_MAIN_FILE)

queue: $(QUEUE_TARGET)

clean-queue:
	-rm $(QUEUE_TARGET)

clean: clean-queue
