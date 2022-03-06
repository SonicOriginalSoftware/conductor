.PHONY: runner clean-runner test-runner

RUNNER_BIN_DIR := bin
RUNNER_MAIN_FILE := $(RUNNER_DIR)/$(RUNNER_BIN_DIR)/main.go
RUNNER_TARGET_FILE := $(RUNNER_DIR)
RUNNER_TARGET := $(OUT_DIR)/$(RUNNER_TARGET_FILE)

$(RUNNER_TARGET): FORCE
	go build -o $(RUNNER_TARGET) $(RUNNER_MAIN_FILE)

runner: $(RUNNER_TARGET)

test-runner:
	@go test -v ./$(RUNNER_DIR)

clean-runner:
	-rm $(RUNNER_TARGET)

clean: clean-runner
