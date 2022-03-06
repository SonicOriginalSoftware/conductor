FORCE:
.DEFAULT_GOAL := help

define USAGE
  conductor
endef

clean:

help:
	$(info $(USAGE))
	@:

OUT_DIR := out
INTEGRATION_TEST_DIR := integration_tests
PROTO_PATH := protos
QUEUE_DIR := queue
RUNNER_DIR := runner

include $(PROTO_PATH)/Makefile
include $(QUEUE_DIR)/Makefile
include $(RUNNER_DIR)/Makefile

include $(INTEGRATION_TEST_DIR)/Makefile

clean-all ca: clean

.PHONY: clean clean-all ca
